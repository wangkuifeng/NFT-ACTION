// script/Deploy.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {console} from "forge-std/console.sol";
import {Script} from "forge-std/Script.sol";
import {NFTAuction} from "../src/NFTAuction.sol";
import {MockERC721} from "../src/MockERC721.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";


contract DeployScript is Script {
    function run() external {
        // 使用 Anvil 的第一个默认私钥
                // 从 .env 里读取私钥
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署 Mock NFT
        MockERC721 nft = new MockERC721();

        // 2. 部署拍卖合约逻辑实现
        NFTAuction implementation = new NFTAuction();

        // 3. 部署代理合约并初始化
        bytes memory data = abi.encodeCall(NFTAuction.initialize, (vm.addr(deployerPrivateKey)));
        ERC1967Proxy proxy = new ERC1967Proxy(address(implementation), data);

        vm.stopBroadcast();

        console.log("====== Deployment Successful ======");
        console.log("Mock NFT Address   :", address(nft));
        console.log("NFTAuction Proxy   :", address(proxy));
    }
}