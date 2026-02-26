'use client';

import { useAccount, useSwitchChain } from 'wagmi';
import { sepolia } from 'wagmi/chains';

export default function NetworkChecker() {
  const { chainId, isConnected } = useAccount();
  const { switchChain } = useSwitchChain();

  // 如果未连接，或者已经连接且在 Sepolia 网络上，则不显示任何东西
  if (!isConnected || chainId === sepolia.id) {
    return null;
  }

  // 如果连错了网络，显示一个无法忽略的红色警告条
  return (
    <div className="fixed top-0 left-0 w-full z-[100] bg-red-500 text-white text-center py-3 font-bold shadow-lg flex items-center justify-center gap-4 animate-pulse">
      <span>⚠️ 警告：你当前连接的不是 Sepolia 测试网，交易将无法进行！</span>
      <button 
        onClick={() => switchChain({ chainId: sepolia.id })}
        className="bg-white text-red-600 px-6 py-1.5 rounded-full text-sm font-black hover:bg-gray-100 hover:scale-105 transition-all shadow-md"
      >
        ⚡️ 一键切换至 Sepolia
      </button>
    </div>
  );
}