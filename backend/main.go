package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"nft-auction-backend/contracts/auction"
	"nft-auction-backend/models"
	// 👉 新增：引入高精度计算包
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

	// 4. [新增] 订阅 AuctionCancelled 事件
	auctionCancelledChan := make(chan *auction.NFTAuctionAuctionCanceled)
	subCanceled, _ := auctionFilterer.WatchAuctionCanceled(nil, auctionCancelledChan, nil)

	fmt.Println("📡 链上事件监听器(全矩阵)已在后台启动...")

	for {
		select {
		case err := <-subCreated.Err():
			log.Printf("Created 订阅出错: %v", err)
		case err := <-subBid.Err():
			log.Printf("Bid 订阅出错: %v", err)
		case err := <-subEnded.Err():
			log.Printf("Ended 订阅出错: %v", err)
		// [新增] 错误处理
		case err := <-subCanceled.Err():
			log.Printf("Canceled 订阅出错: %v", err)
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
		// 5. [新增] 处理拍卖取消
		case event := <-auctionCancelledChan:
			fmt.Println("\n🚫 [事件] 监听到拍卖被取消...")
			db.Model(&models.AuctionRecord{}).
				Where("auction_id = ?", event.AuctionId.String()).
				Updates(map[string]interface{}{
					"status": "Cancelled", // 更新状态为已取消
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

	// API 1: 获取拍卖列表 (已升级：支持状态、卖家、出价者多维过滤)
	r.GET("/api/auctions", func(c *gin.Context) {
		statusFilter := c.Query("status") // active 或 ended
		sellerFilter := c.Query("seller") // 卖家地址
		bidderFilter := c.Query("bidder") // 出价者地址

		var auctions []models.AuctionRecord
		query := db.Model(&models.AuctionRecord{})

		// 1. 卖家过滤 (查询"我创建的")
		if sellerFilter != "" {
			query = query.Where("seller = ?", sellerFilter)
		}

		// 2. 参与者过滤 (查询"我参与的")
		if bidderFilter != "" {
			// 使用子查询：先在 Bid 表中找到该用户出价过的所有 auction_id，再作为主查询条件
			subQuery := db.Model(&models.BidRecord{}).Select("auction_id").Where("bidder = ?", bidderFilter)
			query = query.Where("auction_id IN (?)", subQuery)
		}

		// 3. 状态过滤
		if statusFilter == "ended" {
			// 已结束：包含 Ended 和 Cancelled 状态
			query = query.Where("status IN ?", []string{"Ended", "Cancelled"})
		} else if statusFilter == "active" {
			// 明确指定查进行中
			query = query.Where("status = ?", "Active")
		} else if statusFilter == "" && sellerFilter == "" && bidderFilter == "" {
			// 如果没有任何过滤条件 (通常是首页大厅请求)，默认只显示进行中的
			query = query.Where("status = ?", "Active")
		}
		// 注：如果传了 seller 或 bidder 但没传 status，则返回该用户所有的记录，不限制状态

		// 执行查询并按创建时间倒序
		if err := query.Order("created_at desc").Find(&auctions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询数据库失败"})
			return
		}

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
	// API 4: 平台全局统计数据 (已加入 TVL 计算)
	// ==========================================
	r.GET("/api/stats", func(c *gin.Context) {
		var auctionCount int64
		var bidCount int64

		// 1. 基础统计
		db.Model(&models.AuctionRecord{}).Count(&auctionCount)
		db.Model(&models.BidRecord{}).Count(&bidCount)

		// 2. TVL 计算核心逻辑
		var activeAuctions []models.AuctionRecord
		// 查询所有状态为 "Active" 的拍卖
		db.Where("status = ?", "Active").Find(&activeAuctions)

		totalWei := big.NewInt(0) // 初始化总和为 0

		for _, auc := range activeAuctions {
			// 如果该拍卖有人出价 (最高价不是 0)，则累加进总池子
			if auc.HighestBid != "0" && auc.HighestBid != "" {
				bidAmount, ok := new(big.Int).SetString(auc.HighestBid, 10)
				if ok {
					totalWei.Add(totalWei, bidAmount) // totalWei += bidAmount
				}
			}
		}

		// 3. 将 Wei 转换为 ETH，并格式化为保留 4 位小数的字符串
		tvlEth := new(big.Float).Quo(new(big.Float).SetInt(totalWei), big.NewFloat(1e18))
		tvlString := tvlEth.Text('f', 4)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"total_auctions": auctionCount,
				"total_bids":     bidCount,
				"tvl":            tvlString, // 👉 动态计算的 TVL 输出
			},
		})
	})

	// ==========================================
	// API 5: 集成 Alchemy，获取用户的 NFT 列表
	// ==========================================
	r.GET("/api/users/:address/nfts", func(c *gin.Context) {
		ownerAddress := c.Param("address")

		// // 拦截器：如果是我们本地 Anvil 的测试地址，由于 Alchemy 查不到，直接返回 Mock 数据保障前端开发
		// if strings.EqualFold(ownerAddress, "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266") {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"success": true,
		// 		"data": []gin.H{
		// 			{
		// 				"contract": gin.H{"address": "0x5FbDB2315678afecb367f032d93F642f64180aa3"}, // 你本地部署的 Mock NFT 地址
		// 				"id":       gin.H{"tokenId": "1"},
		// 				"title":    "Mock NFT #1",
		// 				"media":    []gin.H{{"gateway": "https://via.placeholder.com/200"}}, // 假图片占位
		// 			},
		// 		},
		// 	})
		// 	return
		// }

		// 真实的 Alchemy API 调用逻辑 (未来上测试网/主网时生效)
		// 注意：这里的 demo key 仅供测试，未来请去 alchemy.com 申请你自己的 API Key
		alchemyApiKey := "demo"
		url := fmt.Sprintf("https://eth-sepolia.g.alchemy.com/nft/v3/%s/getNFTsForOwner?owner=%s&withMetadata=true", alchemyApiKey, ownerAddress)

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

	// ==========================================
	// API 6: [新增] 用户个人统计数据 (供个人中心调用)
	// ==========================================
	r.GET("/api/users/:address/stats", func(c *gin.Context) {
		address := c.Param("address")
		var createdCount int64
		var participatedCount int64

		// 统计我创建的拍卖数量
		db.Model(&models.AuctionRecord{}).Where("seller = ?", address).Count(&createdCount)

		// 统计我参与的拍卖数量 (Distinct 去重：同一个拍卖出价多次只算1次参与)
		db.Model(&models.BidRecord{}).Where("bidder = ?", address).Distinct("auction_id").Count(&participatedCount)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"created_count":      createdCount,
				"participated_count": participatedCount,
			},
		})
	})

	fmt.Println("🌐 REST API 服务已启动，请访问 http://localhost:8080")
	r.Run("0.0.0.0:8080")
}
