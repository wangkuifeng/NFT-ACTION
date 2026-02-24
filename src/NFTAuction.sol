// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";


contract NFTAuction is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    
    uint256 public nextAuctionId;

    struct Auction {
        address seller;          
        address nftContract;     
        uint256 tokenId;         
        uint256 startPrice;      
        uint256 startTime;       
        uint256 duration;        
        uint256 endTime;         
        address highestBidder;   
        uint256 highestBid;      
        bool active;             
    }

    mapping(address => mapping(uint256 => uint256)) public nftToken2AuctionId; 
    mapping(uint256 => Auction) public auctionData; 

    // ==========================================
    // Events (后端监听数据源) 
    // ==========================================
    event AuctionCreated(
        uint256 indexed auctionId, 
        address indexed seller, 
        address indexed nftContract, 
        uint256 tokenId, 
        uint256 startPrice, 
        uint256 startTime, 
        uint256 endTime
    );
// ==========================================
    // 新增 Events (供后端监听出价记录)
    // ==========================================
    event BidPlaced(
        uint256 indexed auctionId, 
        address indexed bidder, 
        uint256 amount, 
        uint256 timestamp
    );
    event AuctionCanceled(uint256 indexed auctionId);
    event AuctionEnded(
            uint256 indexed auctionId, 
            address indexed winner, 
            uint256 amount
        );
    // ==========================================
    // 新增 Custom Errors
    // ==========================================
    error AuctionNotEnded();
    error NotSeller();
    error AuctionAlreadyStarted();
    // ==========================================
    // 新增 Custom Errors
    // ==========================================
    error AuctionNotActive();
    error AuctionNotStarted();
    error AuctionEndederror();
    error BidTooLow();
    error RefundFailed();
    // ==========================================
    // Custom Errors (自定义错误，比 require 更省 gas)
    // ==========================================
    error InvalidTime();
    error AuctionAlreadyExists();
    error NotNFTOwner();
    // 👇 [新增] 取消时如果已经有人出价，则拦截
    error BidsAlreadyPlaced();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) initializer public {
        __Ownable_init(initialOwner);
        _transferOwnership(initialOwner);
        nextAuctionId = 1; 
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // ==========================================
    // 核心业务逻辑
    // ==========================================

    /**
     * @notice 创建一个新的 NFT 拍卖
     * @param nftContract NFT 合约地址
     * @param tokenId NFT Token ID
     * @param startPrice 起拍价格
     * @param startTime 拍卖开始时间 (如果是0，表示立即开始)
     * @param duration 拍卖持续时间
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startPrice,
        uint256 startTime,
        uint256 duration
    ) external {
        // 1. 校验 (Checks)
        if (duration == 0) revert InvalidTime();
        
        uint256 actualStartTime = startTime == 0 ? block.timestamp : startTime;
        if (actualStartTime < block.timestamp) revert InvalidTime();
        
        // 检查这个 NFT 是否已经在拍卖中
        uint256 existingAuctionId = nftToken2AuctionId[nftContract][tokenId];
        if (existingAuctionId != 0 && auctionData[existingAuctionId].active) {
            revert AuctionAlreadyExists();
        }

        // 检查调用者是否真的拥有这个 NFT
        if (IERC721(nftContract).ownerOf(tokenId) != msg.sender) {
            revert NotNFTOwner();
        }

        // 2. 更新状态 (Effects)
        uint256 auctionId = nextAuctionId;
        uint256 endTime = actualStartTime + duration;

        //上架一个NFT
        auctionData[auctionId] = Auction({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            startPrice: startPrice,
            startTime: actualStartTime,
            duration: duration,
            endTime: endTime,
            highestBidder: address(0),
            highestBid: 0,
            active: true
        });

        nftToken2AuctionId[nftContract][tokenId] = auctionId;
        nextAuctionId++;

        // 触发事件，供 Go 后端抓取存入 MySQL/Redis [cite: 64, 65, 66, 133, 134]
        emit AuctionCreated(auctionId, msg.sender, nftContract, tokenId, startPrice, actualStartTime, endTime);

        // 3. 交互 (Interactions)
        // 将 NFT 从卖家转移到本合约进行锁定 
        // 注意：在此调用之前，前端必须提示用户调用 NFT 合约的 approve() 或 setApprovalForAll() 授权给本合约
        IERC721(nftContract).transferFrom(msg.sender, address(this), tokenId);
    }

    /**
     * @notice 买家参与出价
     * @param auctionId 拍卖的 ID
     */
    function bidAuction(uint256 auctionId) external payable {
        // 使用 storage 引用直接操作合约里的状态数据
        Auction storage auction = auctionData[auctionId];

        // 1. 校验 (Checks)
        if (!auction.active) revert AuctionNotActive();
        if (block.timestamp < auction.startTime) revert AuctionNotStarted();
        if (block.timestamp >= auction.endTime) revert AuctionEndederror();
        
        // 判断出价是否足够高：如果是首次出价，必须大于等于起拍价；否则必须大于当前最高价
        if (auction.highestBid == 0) {
            // 首次出价：必须大于等于起拍价
            if (msg.value < auction.startPrice) revert BidTooLow();
        } else {
            // 后续抢拍：必须严格大于当前最高价
            if (msg.value <= auction.highestBid) revert BidTooLow();
        }

        // 记录前一个最高出价者和金额，用于稍后的退款 
        address previousBidder = auction.highestBidder;
        uint256 previousBidAmount = auction.highestBid;

        // 2. 更新状态 (Effects) - 【非常重要：必须在退款前修改状态】
        auction.highestBidder = msg.sender;
        auction.highestBid = msg.value;

        // 触发事件，Go 后端监听到这个事件后，会将出价记录写入 MySQL [cite: 133, 134]
        emit BidPlaced(auctionId, msg.sender, msg.value, block.timestamp);

        // 3. 交互 (Interactions) - 退款给前一个买家 
        if (previousBidder != address(0)) {
            // 使用 call 进行底层的以太坊转账，比 transfer 更安全、兼容性更好
            (bool success, ) = payable(previousBidder).call{value: previousBidAmount}("");
            if (!success) revert RefundFailed();
        }
    }

    /**
     * @notice 结束拍卖 (交割)
     * @param auctionId 拍卖的 ID
     */
    function endAuction(uint256 auctionId) external {
        Auction storage auction = auctionData[auctionId];

        // 1. 校验 (Checks)
        if (!auction.active) revert AuctionNotActive();
        if (block.timestamp < auction.endTime) revert AuctionNotEnded();

        // 2. 更新状态 (Effects)
        // 将拍卖标记为已结束
        auction.active = false;
        // 清除 NFT 与拍卖 ID 的映射，以便该 NFT 以后还能再次被拍卖
        delete nftToken2AuctionId[auction.nftContract][auction.tokenId];

        address winner = auction.highestBidder;
        uint256 amount = auction.highestBid;
        address seller = auction.seller;

        // 触发事件，Go 后端监听到后会更新数据库中的拍卖状态
        emit AuctionEnded(auctionId, winner, amount);

        // 3. 交互 (Interactions)
        if (winner != address(0)) {
            // 情况 A：有人出价获胜
            // 买家拿走 NFT 
            IERC721(auction.nftContract).transferFrom(address(this), winner, auction.tokenId);
            // 资金转给卖家 
            (bool success, ) = payable(seller).call{value: amount}("");
            if (!success) revert RefundFailed();
        } else {
            // 情况 B：流拍 (没人参与出价)
            // 原来的持有者拿回自己的 NFT 
            IERC721(auction.nftContract).transferFrom(address(this), seller, auction.tokenId);
        }
    }

    /**
     * @notice 取消拍卖 (卖家反悔)
     * @param auctionId 拍卖的 ID
     */
    function cancelAuction(uint256 auctionId) external {
        Auction storage auction = auctionData[auctionId];

        // 1. 校验 (Checks)
        if (!auction.active) revert AuctionNotActive();
        // 只有卖家自己可以取消
        if (msg.sender != auction.seller) revert NotSeller();
        
        // 👇👇 【核心修改点】：不再校验开始时间，而是校验是否有人出价
        if (auction.highestBidder != address(0)) revert BidsAlreadyPlaced();

        // 2. 更新状态 (Effects)
        auction.active = false;
        delete nftToken2AuctionId[auction.nftContract][auction.tokenId];

        emit AuctionCanceled(auctionId);

        // 3. 交互 (Interactions)
        // 将 NFT 退还给卖家
        IERC721(auction.nftContract).transferFrom(address(this), auction.seller, auction.tokenId);
    }
}