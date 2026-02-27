import { http, createConfig } from 'wagmi'
//import { mainnet, localhost } from 'wagmi/chains'
import { mainnet, sepolia } from 'wagmi/chains'
import { injected } from 'wagmi/connectors'

export const config = createConfig({
  chains: [mainnet, sepolia], 
  connectors: [
    injected(), // 自动识别 MetaMask 等插件钱包
  ],
  transports: {
    [mainnet.id]: http(),
    // 这里的 http 必须对应你运行 anvil 的地址
    [sepolia.id]: http('https://sepolia.infura.io/v3/a4583e6142214f4a8eb27d048a906649'), 
  },
})