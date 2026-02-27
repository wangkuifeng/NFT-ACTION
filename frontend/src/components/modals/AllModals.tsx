// src/components/modals/AllModals.tsx
import React from 'react';

// ==========================================
// 1. 极速上架向导弹窗 (CreateAuctionModal)
// ==========================================
export function CreateAuctionModal({
  isOpen, onClose, 
  formNftContract, setFormNftContract,
  formTokenId, setFormTokenId,
  formStartPrice, setFormStartPrice,
  formDuration, setFormDuration,
  handleApproveAndList, isListing
}: any) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-3xl p-8 max-w-md w-full shadow-2xl">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-black">上架向导</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-red-500 font-bold">✕</button>
        </div>
        
        <div className="space-y-4">
          <div>
            <label className="block text-xs font-bold text-gray-500 mb-1">NFT 合约地址 (已自动填入)</label>
            <input value={formNftContract} onChange={e => setFormNftContract(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-xs font-bold text-gray-500 mb-1">Token ID</label>
              <input type="number" placeholder="点右上角领装备" value={formTokenId} onChange={e => setFormTokenId(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
            </div>
            <div>
              <label className="block text-xs font-bold text-gray-500 mb-1">起拍价 (ETH)</label>
              <input type="number" step="0.01" value={formStartPrice} onChange={e => setFormStartPrice(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
            </div>
          </div>
          <div>
            <label className="block text-xs font-bold text-gray-500 mb-1">拍卖时长 (天)</label>
            <input type="number" value={formDuration} onChange={e => setFormDuration(e.target.value)} className="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-2 text-sm focus:outline-blue-500" />
          </div>

          <div className="mt-8">
            <button 
              onClick={handleApproveAndList} 
              disabled={isListing}
              className="w-full bg-blue-600 text-white py-4 rounded-xl font-bold shadow-lg shadow-blue-200 hover:bg-blue-700 transition-all duration-300 disabled:opacity-50 text-lg flex flex-col items-center justify-center"
            >
              <span>{isListing ? '⏳ 链上处理中... (需连续确认两次交易)' : '一键授权并上架'}</span>
              {!isListing && <span className="text-[10px] font-normal mt-1 opacity-80">小狐狸将弹出两次，请耐心等待区块确认</span>}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

// ==========================================
// 2. 我的 NFT 画廊弹窗 (NFTGalleryModal)
// ==========================================
export function NFTGalleryModal({
  isOpen, onClose, isLoadingNFTs, myNFTs, handleSelectNFTForAuction
}: any) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-3xl p-8 max-w-4xl w-full shadow-2xl max-h-[85vh] flex flex-col">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-black">我的 NFT 资产库</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-red-500 font-bold text-xl">✕</button>
        </div>
        
        <div className="overflow-y-auto pr-2 custom-scrollbar">
          {isLoadingNFTs ? (
            <div className="flex flex-col items-center justify-center py-20">
              <span className="text-4xl animate-bounce mb-4">📡</span>
              <p className="text-gray-500 font-bold">正在从节点同步你的资产...</p>
            </div>
          ) : myNFTs.length > 0 ? (
            <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
              {myNFTs.map((nft: any, idx: number) => (
                <div key={idx} className="bg-white border rounded-2xl overflow-hidden shadow-sm hover:shadow-xl transition-all duration-300 group relative">
                  <div className="aspect-square bg-gray-50 flex items-center justify-center relative">
                    {nft.imageUrl || nft.image_url ? (
                      <img 
                        src={nft.imageUrl || nft.image_url} 
                        alt={nft.name} 
                        className="object-cover w-full h-full" 
                        onError={(e) => { e.currentTarget.src = 'https://via.placeholder.com/150?text=No+Image' }} 
                      />
                    ) : (
                      <span className="text-4xl">🖼️</span>
                    )}
                    <div className="absolute inset-0 bg-black/60 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center backdrop-blur-[1px]">
                      <button 
                        onClick={() => handleSelectNFTForAuction(nft)}
                        className="bg-blue-600 text-white px-5 py-2.5 rounded-xl font-bold shadow-lg hover:scale-105 transition-transform"
                      >
                        ⚡️ 一键上架
                      </button>
                    </div>
                  </div>
                  <div className="p-4">
                    <p className="text-[10px] text-gray-400 font-bold uppercase truncate mb-1">
                      {nft.collectionName || nft.collection_name || 'Unknown Collection'}
                    </p>
                    <h3 className="text-sm font-black truncate text-slate-800">
                      {nft.name || `Token #${nft.tokenId || nft.token_id}`}
                    </h3>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="text-center py-20 bg-gray-50 rounded-3xl border-2 border-dashed border-gray-200">
              <span className="text-4xl mb-4 block">🍃</span>
              <p className="text-gray-500 font-bold">你的钱包里还没有 NFT 哦</p>
              <p className="text-xs text-gray-400 mt-2">去领一个测试 NFT 试试看吧</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

// ==========================================
// 3. 拍卖详情与出价历史弹窗 (AuctionDetailsModal)
// ==========================================
export function AuctionDetailsModal({
  isOpen, onClose, selectedItem, isLoadingBids, bidHistory, openBidModal
}: any) {
  if (!isOpen || !selectedItem) return null;

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-3xl p-8 max-w-2xl w-full shadow-2xl flex flex-col max-h-[90vh]">
        <div className="flex justify-between items-center mb-6 border-b pb-4">
          <h2 className="text-2xl font-black text-slate-900">拍卖详情: Token #{selectedItem.token_id}</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-red-500 font-bold text-xl">✕</button>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
          <div className="bg-gray-50 rounded-2xl p-6 flex flex-col justify-center items-center text-center">
            <div className="text-6xl mb-4">🖼️</div>
            <p className="text-gray-500 text-xs mb-1">合约地址</p>
            <p className="font-mono text-sm truncate w-full mb-4 px-4">{selectedItem.nft_contract}</p>
            <p className="text-gray-500 text-xs mb-1">当前最高价</p>
            <p className="text-3xl font-black text-blue-600 mb-6">
              {selectedItem.highest_bid === "0" ? (Number(selectedItem.start_price) / 1e18).toFixed(4) : (Number(selectedItem.highest_bid) / 1e18).toFixed(4)} ETH
            </p>
            
            {/* 传递事件：先关掉自己，再打开竞拍弹窗 */}
            {selectedItem.status !== 'Ended' && Date.now() < Number(selectedItem.end_time) * 1000 && (
              <button 
                onClick={() => {
                  onClose(); 
                  openBidModal(selectedItem);
                }} 
                className="w-full bg-blue-600 text-white py-3 rounded-xl font-bold shadow-lg shadow-blue-200 hover:bg-blue-700 transition"
              >
                我要出价 🔨
              </button>
            )}
          </div>

          <div className="flex flex-col">
            <h3 className="text-lg font-bold mb-4 flex items-center gap-2">
              <span>📜 出价历史</span>
              <span className="bg-blue-100 text-blue-600 text-xs px-2 py-0.5 rounded-full">{bidHistory.length} 次</span>
            </h3>
            
           <div className="flex-1 overflow-y-auto max-h-75 pr-2 custom-scrollbar bg-white rounded-xl border border-gray-100">
              {isLoadingBids ? (
                <div className="text-center py-10 text-gray-400">加载流水记录中...</div>
              ) : bidHistory.length > 0 ? (
                <div className="divide-y divide-gray-100">
                    {bidHistory.map((bid: any, index: number) => {
                      const bidder = bid.Bidder || bid.bidder || bid.bidder_address || "";
                      const amount = bid.Amount || bid.amount || bid.bid_amount || "0";
                      const timestamp = bid.Timestamp || bid.timestamp || bid.created_at || (Date.now() / 1000);
                      return (
                        <div key={index} className="p-3 hover:bg-gray-50 transition flex justify-between items-center text-sm">
                          <div>
                            <p className="font-mono font-bold text-slate-700">
                              {bidder ? `${bidder.slice(0, 6)}...${bidder.slice(-4)}` : '未知出价者'}
                            </p>
                            <p className="text-[10px] text-gray-400 mt-1">
                              {new Date(timestamp * 1000).toLocaleString()}
                            </p>
                          </div>
                          <div className="font-black text-blue-600">
                            {(Number(amount) / 1e18).toFixed(4)} ETH
                          </div>
                        </div>
                      );
                    })}
                </div>
              ) : (
                <div className="text-center py-12 text-gray-400 flex flex-col items-center">
                  <span className="text-2xl mb-2">👻</span>
                  <p>暂时还没有人出价</p>
                  <p className="text-xs mt-1">抢个首发沙发吧！</p>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}