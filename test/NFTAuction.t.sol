// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {NFTAuction} from "../src/NFTAuction.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
// ==========================================
// 1. 搞一个假的 NFT 合约用来测试
// ==========================================
contract MockERC721 is ERC721 {
    constructor() ERC721("Mock NFT", "MNFT") {}
    
    // 开放一个 mint 函数给测试用
    function mint(address to, uint256 tokenId) external {
        _mint(to, tokenId);
    }
}

// ==========================================
// 2. 核心测试合约
// ==========================================
contract NFTAuctionTest is Test {
    NFTAuction public auction;
    MockERC721 public nft;

    // 定义几个测试用的钱包地址
    address public seller = address(0x1);
    address public buyer1 = address(0x2);
    address public buyer2 = address(0x3);

    function setUp() public {
            // 1. 先部署“逻辑合约” (Implementation)
            NFTAuction implementation = new NFTAuction();

            // 2. 将我们要调用的 initialize 函数及其参数打包成 calldata
            bytes memory data = abi.encodeCall(NFTAuction.initialize, (address(this)));

            // 3. 部署“代理合约” (Proxy)，将它指向逻辑合约，并在部署时立即执行 initialize
            ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), data);

            // 4. 让我们的 auction 变量指向代理合约的地址 (这样后续所有的调用都会经过代理)
            auction = NFTAuction(address(proxy));

            // --- 下面是原来的 Mock NFT 和发钱逻辑 ---
            nft = new MockERC721();

            vm.deal(seller, 10 ether);
            vm.deal(buyer1, 10 ether);
            vm.deal(buyer2, 10 ether);

            nft.mint(seller, 1);
        }

    // 核心测试：模拟完整的拍卖生命周期
    function testFullAuctionFlow() public {
        // ==========================
        // 步骤 A: 卖家创建拍卖
        // ==========================
        // vm.startPrank 会将接下来所有交易的发起者(msg.sender)伪装成 seller
        vm.startPrank(seller);
        
        // 卖家必须先授权给拍卖合约
        nft.approve(address(auction), 1); 
        
        uint256 startPrice = 1 ether;
        uint256 startTime = block.timestamp; // 立即开始
        uint256 duration = 1 days;           // 拍卖持续 1 天

        auction.createAuction(address(nft), 1, startPrice, startTime, duration);
        vm.stopPrank();

        // 断言(Assert)：验证 NFT 确实从卖家转移到了拍卖合约中被锁定
        assertEq(nft.ownerOf(1), address(auction));

        // ==========================
        // 步骤 B: 买家 1 出价
        // ==========================
        vm.startPrank(buyer1);
        auction.bidAuction{value: 1.5 ether}(1); // 传入拍卖 ID: 1
        vm.stopPrank();

        // ==========================
        // 步骤 C: 买家 2 抢拍 (验证退款逻辑)
        // ==========================
        uint256 buyer1BalanceBefore = buyer1.balance; // 记录买家1被抢拍前的余额
        
        vm.startPrank(buyer2);
        auction.bidAuction{value: 2 ether}(1); // 买家2 出价 2 ETH
        vm.stopPrank();

        // 断言：买家1 的余额应该增加了 1.5 ETH (说明退款成功)
        uint256 buyer1BalanceAfter = buyer1.balance;
        assertEq(buyer1BalanceAfter - buyer1BalanceBefore, 1.5 ether);

        // ==========================
        // 步骤 D: 时间快进 (模拟拍卖到期)
        // ==========================
        // vm.warp 是时间魔法，直接把区块链时间快进 1天零1秒
        vm.warp(block.timestamp + 1 days + 1 seconds);

        // ==========================
        // 步骤 E: 结束拍卖 (交割)
        // ==========================
        uint256 sellerBalanceBefore = seller.balance;

        // 任何人都可以触发结束拍卖（这里模拟买家2自己去触发）
        vm.startPrank(buyer2);
        auction.endAuction(1);
        vm.stopPrank();

        // 断言：买家2 成功拿到了 NFT
        assertEq(nft.ownerOf(1), buyer2);

        // 断言：卖家 成功拿到了 2 ETH 的货款
        uint256 sellerBalanceAfter = seller.balance;
        assertEq(sellerBalanceAfter - sellerBalanceBefore, 2 ether);
    }

    // ==========================================
    // 边界与覆盖测试 (Boundary & Coverage Tests)
    // ==========================================

    // 1. 边界测试：创建拍卖时，时间参数非法
    function testRevert_CreateAuction_InvalidTime() public {
        vm.startPrank(seller);
        nft.approve(address(auction), 1);

        // 边界：测试 duration 为 0
        vm.expectRevert(NFTAuction.InvalidTime.selector);
        auction.createAuction(address(nft), 1, 1 ether, block.timestamp, 0);

        // --- 新增这一行：把时间往后拨，确保 block.timestamp 足够大 ---
        vm.warp(1000); 

        // 边界：此时 block.timestamp 是 1000，传入 999 就会被正常拦截
        vm.expectRevert(NFTAuction.InvalidTime.selector);
        auction.createAuction(address(nft), 1, 1 ether, block.timestamp - 1, 1 days);
        vm.stopPrank();
    }

    // 2. 权限覆盖测试：不是 NFT 的拥有者尝试发起拍卖
    function testRevert_CreateAuction_NotOwner() public {
        vm.startPrank(buyer1); // buyer1 并没有 tokenId 为 1 的 NFT
        vm.expectRevert(NFTAuction.NotNFTOwner.selector);
        auction.createAuction(address(nft), 1, 1 ether, block.timestamp, 1 days);
        vm.stopPrank();
    }

    // 3. 金额边界测试：出价金额等于或小于当前最高价/起拍价
    function testRevert_BidAuction_BidTooLow() public {
        vm.startPrank(seller);
        nft.approve(address(auction), 1);
        auction.createAuction(address(nft), 1, 1 ether, block.timestamp, 1 days);
        vm.stopPrank();

        vm.startPrank(buyer1);
        // 边界：首次出价刚好低于起拍价
        vm.expectRevert(NFTAuction.BidTooLow.selector);
        auction.bidAuction{value: 0.99 ether}(1);

        // 正常出价 (压线出价 1 ether 是允许的，因为起拍价是 1 ether)
        auction.bidAuction{value: 1 ether}(1);
        vm.stopPrank();

        vm.startPrank(buyer2);
        // 边界：二次出价刚好等于当前最高价 (1 ether)，这应该被拒绝
        vm.expectRevert(NFTAuction.BidTooLow.selector);
        auction.bidAuction{value: 1 ether}(1);
        vm.stopPrank();
    }

    // 4. 时间边界测试：在未开始或已结束后尝试出价
    function testRevert_BidAuction_TimeBoundaries() public {
        vm.startPrank(seller);
        nft.approve(address(auction), 1);
        // 设定 1 天后才正式开始
        uint256 startTime = block.timestamp + 1 days;
        auction.createAuction(address(nft), 1, 1 ether, startTime, 1 days);
        vm.stopPrank();

        vm.startPrank(buyer1);
        // 边界：还没开始就抢跑出价
        vm.expectRevert(NFTAuction.AuctionNotStarted.selector);
        auction.bidAuction{value: 2 ether}(1);
        vm.stopPrank();

        // 魔法快进：将时间直接拨到拍卖结束之后
        vm.warp(startTime + 1 days + 1 seconds);
        
        vm.startPrank(buyer2);
        // 边界：结束后试图出价
        vm.expectRevert(abi.encodeWithSelector(NFTAuction.AuctionEndederror.selector));
        auction.bidAuction{value: 2 ether}(1);

        vm.stopPrank();
    }

    // 5. 状态与权限边界测试：取消拍卖 (CancelAuction) 逻辑
    function test_CancelAuction_Boundaries() public {
        vm.startPrank(seller);
        nft.approve(address(auction), 1);
        // 设定 1 天后才正式开始，这样才有"反悔期"
        uint256 startTime = block.timestamp + 1 days;
        auction.createAuction(address(nft), 1, 1 ether, startTime, 1 days);
        vm.stopPrank();

        // 权限越界：买家试图取消卖家的拍卖
        vm.startPrank(buyer1);
        vm.expectRevert(NFTAuction.NotSeller.selector);
        auction.cancelAuction(1);
        vm.stopPrank();

        // 正常取消流程
        vm.startPrank(seller);
        auction.cancelAuction(1);
        
        // 断言：验证取消后，NFT 确实退回给了卖家
        assertEq(nft.ownerOf(1), seller);
        vm.stopPrank();
    }

    // 6. 状态覆盖测试：不能对未激活或已结束的拍卖执行交割
    function testRevert_EndAuction_NotActiveOrNotEnded() public {
        vm.startPrank(seller);
        nft.approve(address(auction), 1);
        auction.createAuction(address(nft), 1, 1 ether, block.timestamp, 1 days);
        vm.stopPrank();

        // 边界：时间还没到就想提前结束拿钱
        vm.expectRevert(NFTAuction.AuctionNotEnded.selector);
        auction.endAuction(1);

        // 快进到结束
        vm.warp(block.timestamp + 1 days + 1 seconds);
        auction.endAuction(1); // 第一次交割成功

        // 边界：已经交割过的拍卖，再次尝试交割
        vm.expectRevert(NFTAuction.AuctionNotActive.selector);
        auction.endAuction(1); 
    }
}