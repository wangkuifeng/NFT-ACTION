// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {NFTAuction} from "../src/NFTAuction.sol";
import {console} from "forge-std/console.sol";

contract UpgradeAuction is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // 这是你之前部署好、并且已经在前端配置好的那个 Proxy 地址
        address proxyAddress = 0xbf1304F2D77D110a32B150792Ee481212926763D; 

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署全新的逻辑合约 (Implementation)
        NFTAuction newImplementation = new NFTAuction();
        
        // 2. 调用 Proxy 的 upgradeToAndCall 方法进行升级
        // 注意：因为 UUPS 的升级逻辑写在实现合约里，所以我们把 Proxy 地址包装成 NFTAuction 接口来调用
        NFTAuction proxy = NFTAuction(proxyAddress);
        proxy.upgradeToAndCall(address(newImplementation), "");

        vm.stopBroadcast();
        
        // 打印新逻辑合约地址（仅供参考，前端依然只认 proxyAddress）
        console.log("New Implementation deployed to:", address(newImplementation));
        console.log("Proxy successfully upgraded!");
    }
}