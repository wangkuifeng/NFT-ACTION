package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"nft-auction-backend/contracts/auction"
	"nft-auction-backend/models"
)

func initDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/nft_auction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	// 自动迁移两张表
	db.AutoMigrate(&models.AuctionRecord{}, &models.BidRecord{})
	return db
}

func startEventListener(db *gorm.DB) {
	//nodeURL := "ws://127.0.0.1:8545"
	nodeURL := "wss://sepolia.infura.io/ws/v3/a4583e6142214f4a8eb27d048a906649"
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("无法连接到以太坊节点: %v", err)
	}
	defer client.Close()

	// ⚠️ 记得换成你本地部署的真实 Proxy 合约地址
	contractAddress := common.HexToAddress("0xbf1304F2D77D110a32B150792Ee481212926763D")
	auctionFilterer, err := auction.NewNFTAuctionFilterer(contractAddress, client)
	if err != nil {
		log.Fatalf("实例化合约过滤器失败: %v", err)
	}

	// 1. 订阅 AuctionCreated 事件 [cite: 133, 134]
	auctionCreatedChan := make(chan *auction.NFTAuctionAuctionCreated)
	subCreated, _ := auctionFilterer.WatchAuctionCreated(nil, auctionCreatedChan, nil, nil, nil)

	// 2. 订阅 BidPlaced 事件 [cite: 133, 134]
	bidPlacedChan := make(chan *auction.NFTAuctionBidPlaced)
	subBid, _ := auctionFilterer.WatchBidPlaced(nil, bidPlacedChan, nil, nil)

	// 3. 订阅 AuctionEnded 事件 [cite: 133, 134]
	auctionEndedChan := make(chan *auction.NFTAuctionAuctionEnded)
	subEnded, _ := auctionFilterer.WatchAuctionEnded(nil, auctionEndedChan, nil, nil)

	fmt.Println("📡 链上事件监听器(全矩阵)已在后台启动...")

	for {
		select {
		case err := <-subCreated.Err():
			log.Printf("Created 订阅出错: %v", err)
		case err := <-subBid.Err():
			log.Printf("Bid 订阅出错: %v", err)
		case err := <-subEnded.Err():
			log.Printf("Ended 订阅出错: %v", err)

		// 处理创建拍卖
		case event := <-auctionCreatedChan:
			fmt.Println("\n🔥 [事件] 监听到创建拍卖...")
			newRecord := models.AuctionRecord{
				AuctionID:   event.AuctionId.String(),
				Seller:      event.Seller.Hex(),
				NFTContract: event.NftContract.Hex(),
				TokenID:     event.TokenId.String(),
				StartPrice:  event.StartPrice.String(),
				StartTime:   event.StartTime.Uint64(),
				EndTime:     event.EndTime.Uint64(),
				Status:      "Active",
			}
			db.Create(&newRecord)

		// 处理用户出价
		case event := <-bidPlacedChan:
			fmt.Println("\n💰 [事件] 监听到新的出价...")
			// A. 记录到出价历史表
			newBid := models.BidRecord{
				AuctionID: event.AuctionId.String(),
				Bidder:    event.Bidder.Hex(),
				Amount:    event.Amount.String(),
				Timestamp: event.Timestamp.Uint64(),
			}
			db.Create(&newBid)

			// B. 更新主拍卖表的当前最高价和最高出价人
			db.Model(&models.AuctionRecord{}).
				Where("auction_id = ?", event.AuctionId.String()).
				Updates(map[string]interface{}{
					"highest_bid":    event.Amount.String(),
					"highest_bidder": event.Bidder.Hex(),
				})

		// 处理交割结束
		case event := <-auctionEndedChan:
			fmt.Println("\n🏁 [事件] 监听到拍卖结束交割...")
			db.Model(&models.AuctionRecord{}).
				Where("auction_id = ?", event.AuctionId.String()).
				Updates(map[string]interface{}{
					"status": "Ended",
					"winner": event.Winner.Hex(),
				})
		}
	}
}

func main() {
	db := initDB()
	go startEventListener(db)

	r := gin.Default()

	// CORS 跨域配置
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API 1: 获取所有进行中的拍卖列表 [cite: 129]
	r.GET("/api/auctions", func(c *gin.Context) {
		var auctions []models.AuctionRecord
		db.Where("status = ?", "Active").Find(&auctions)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": auctions})
	})

	// API 2: 获取某个拍卖的详情
	r.GET("/api/auctions/:id", func(c *gin.Context) {
		id := c.Param("id")
		var auction models.AuctionRecord
		if err := db.Where("auction_id = ?", id).First(&auction).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "找不到该拍卖"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": auction})
	})

	// API 3: 获取某个拍卖的出价历史记录
	r.GET("/api/auctions/:id/bids", func(c *gin.Context) {
		id := c.Param("id")
		var bids []models.BidRecord
		// 按时间倒序排列出价记录
		db.Where("auction_id = ?", id).Order("timestamp desc").Find(&bids)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": bids})
	})

	// ==========================================
	// API 4: 平台统计数据 (供首页调用)
	// ==========================================
	r.GET("/api/stats", func(c *gin.Context) {
		var auctionCount int64
		var bidCount int64

		// 使用 GORM 统计数据库中表的总行数
		db.Model(&models.AuctionRecord{}).Count(&auctionCount)
		db.Model(&models.BidRecord{}).Count(&bidCount)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"total_auctions": auctionCount,
				"total_bids":     bidCount,
				"tvl":            "0", // 进阶需求：计算所有活跃拍卖的最高价总和，这里暂用 0 占位
			},
		})
	})

	// ==========================================
	// API 5: 集成 Alchemy，获取用户的 NFT 列表
	// ==========================================
	r.GET("/api/users/:address/nfts", func(c *gin.Context) {
		ownerAddress := c.Param("address")

		// 拦截器：如果是我们本地 Anvil 的测试地址，由于 Alchemy 查不到，直接返回 Mock 数据保障前端开发
		if strings.EqualFold(ownerAddress, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266") {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data": []gin.H{
					{
						"contract": gin.H{"address": "0x5FbDB2315678afecb367f032d93F642f64180aa3"}, // 你本地部署的 Mock NFT 地址
						"id":       gin.H{"tokenId": "1"},
						"title":    "Mock NFT #1",
						"media":    []gin.H{{"gateway": "https://via.placeholder.com/200"}}, // 假图片占位
					},
				},
			})
			return
		}

		// 真实的 Alchemy API 调用逻辑 (未来上测试网/主网时生效)
		// 注意：这里的 demo key 仅供测试，未来请去 alchemy.com 申请你自己的 API Key
		alchemyApiKey := "demo"
		url := fmt.Sprintf("https://eth-mainnet.g.alchemy.com/nft/v3/%s/getNFTsForOwner?owner=%s&withMetadata=true", alchemyApiKey, ownerAddress)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "请求 Alchemy API 失败"})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var alchemyResult map[string]interface{}
		json.Unmarshal(body, &alchemyResult)

		// 将 Alchemy 返回的 ownedNfts 数组直接透传给前端
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    alchemyResult["ownedNfts"],
		})
	})

	fmt.Println("🌐 REST API 服务已启动，请访问 http://localhost:8080")
	r.Run("0.0.0.0:8080")
}
