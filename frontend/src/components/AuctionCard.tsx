// src/components/AuctionCard.tsx
import React, { useState } from 'react';
import { useWriteContract, usePublicClient } from 'wagmi';
import { AUCTION_CONTRACT_ADDRESS, AUCTION_ABI } from '@/constants/contracts';

export default function AuctionCard({ item, currentUserAddress, onRefresh, onOpenDetails }: any) {
  const [isProcessing, setIsProcessing] = useState<'cancel' | 'end' | null>(null);
  
  const publicClient = usePublicClient();
  const { writeContractAsync } = useWriteContract();

  // 状态计算
  const isSeller = currentUserAddress?.toLowerCase() === item.seller?.toLowerCase();
  const hasNoBids = item.highest_bid === "0";
  const isExpired = Date.now() > Number(item.end_time) * 1000;
  const isEnded = ['Ended', 'Cancelled', 'ended', 'cancelled'].includes(item.status || item.Status);

  // 取消拍卖
  const handleCancel = async () => {
    try {
      setIsProcessing('cancel');
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'cancelAuction',
        args: [BigInt(item.auction_id)],
      });
      await publicClient!.waitForTransactionReceipt({ hash });
      setTimeout(() => onRefresh(), 1500);
    } catch (err) {
      console.error("取消失败", err);
    } finally {
      setIsProcessing(null);
    }
  };

  // 结束拍卖
  const handleEnd = async () => {
    try {
      setIsProcessing('end');
      const hash = await writeContractAsync({
        address: AUCTION_CONTRACT_ADDRESS,
        abi: AUCTION_ABI,
        functionName: 'endAuction',
        args: [BigInt(item.auction_id)],
      });
      await publicClient!.waitForTransactionReceipt({ hash });
      setTimeout(() => onRefresh(), 1500);
    } catch (err) {
      console.error("结束失败", err);
    } finally {
      setIsProcessing(null);
    }
  };

  return (
    <div className="bg-white rounded-[2.5rem] p-4 shadow-xl shadow-gray-200/50 hover:-translate-y-2 transition-all duration-300 group">
      <div className="aspect-square bg-linear-to-tr from-blue-50 to-purple-50 rounded-4xl mb-4 flex items-center justify-center text-6xl group-hover:scale-95 transition-transform relative">
        🖼️
        {isExpired && (
          <div className="absolute inset-0 bg-white/60 backdrop-blur-[2px] rounded-4xl flex items-center justify-center">
            <span className="bg-red-500 text-white px-4 py-2 rounded-xl font-bold rotate-[-10deg] shadow-lg">已结束</span>
          </div>
        )}
      </div>
      
      <div className="px-2">
        <div className="flex justify-between items-center mb-2">
          <h3 className="text-xl font-bold">Token #{item.token_id}</h3>
          {!isExpired && <span className="text-[10px] bg-blue-600 text-white px-2 py-0.5 rounded-full font-bold">LIVE</span>}
        </div>
        <p className="text-gray-400 text-xs mb-6 truncate" title={item.nft_contract}>合约: {item.nft_contract}</p>
        
        <div className="flex justify-between items-center bg-gray-50 p-4 rounded-2xl">
          <div className="min-w-0 pr-2">
            <p className="text-[10px] text-gray-400 font-bold uppercase">当前最高价</p>
            <p className={`text-xl font-black truncate ${isExpired ? 'text-gray-500' : 'text-blue-600'}`}>
              {hasNoBids ? (Number(item.start_price) / 1e18).toFixed(4) : (Number(item.highest_bid) / 1e18).toFixed(4)} ETH
            </p>
          </div>
          
          <div className="flex gap-2 shrink-0">
            {isEnded ? (
              <button disabled className="bg-gray-100 text-gray-400 border border-gray-200 px-5 py-2.5 rounded-xl text-sm font-bold cursor-not-allowed whitespace-nowrap">
                {/cancel/i.test(item.status || item.Status || '') ? '已取消 🚫' : '已完结 🏁'}
              </button>
            ) : (
              <>
                {isSeller && hasNoBids && !isExpired && (
                  <button onClick={handleCancel} disabled={isProcessing === 'cancel'} className="bg-red-50 text-red-600 border border-red-200 px-4 py-2.5 rounded-xl text-sm font-bold hover:bg-red-100 transition-colors disabled:opacity-50 whitespace-nowrap">
                    {isProcessing === 'cancel' ? '...' : '取消'}
                  </button>
                )}
                {isExpired ? (
                  <button onClick={handleEnd} disabled={isProcessing === 'end'} className="bg-green-500 text-white px-5 py-2.5 rounded-xl text-sm font-bold hover:bg-green-600 transition-colors shadow-md disabled:opacity-50 whitespace-nowrap">
                    {isProcessing === 'end' ? '结算中...' : '结束交割'}
                  </button>
                ) : (
                  <button onClick={() => onOpenDetails(item)} className="bg-slate-900 text-white px-5 py-2.5 rounded-xl text-sm font-bold hover:bg-blue-600 transition-colors whitespace-nowrap">
                    看详情
                  </button>
                )}
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}