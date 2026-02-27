// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {NFTAuction} from "../src/NFTAuction.sol";
import {console} from "forge-std/console.sol";

contract UpgradeAuction is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        // 获取部署者的钱包地址，准备作为默认的平台金库
        address deployerAddress = vm.addr(deployerPrivateKey); 
        
        // 你的 Proxy 地址 (保持不变)
        address proxyAddress = 0xbf1304F2D77D110a32B150792Ee481212926763D; 

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署全新的逻辑合约 (Implementation)
        NFTAuction newImplementation = new NFTAuction();
        
        // 2. 获取 Proxy 实例
        NFTAuction proxy = NFTAuction(proxyAddress);

        // 3. 执行升级：将 Proxy 的逻辑指针指向新合约
        proxy.upgradeToAndCall(address(newImplementation), "");

        // ==========================================
        // 4. 👇 【关键新增】初始化新增的状态变量！
        // ==========================================
        // 直接调用新版合约中的 setter 函数，为平台抽成和金库地址赋值
        proxy.setPlatformFee(200); // 设置 2% 抽成 (200 BPS)
        proxy.setFeeRecipient(deployerAddress); // 默认将抽成打入部署者的钱包

        vm.stopBroadcast();
        
        // 打印执行结果
        console.log("New Implementation deployed to:", address(newImplementation));
        console.log("Proxy successfully upgraded!");
        console.log("Platform Fee successfully set to 2% (200 BPS)");
        console.log("Fee Recipient set to:", deployerAddress);
    }
}