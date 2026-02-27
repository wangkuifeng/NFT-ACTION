// src/services/api.ts

const API_BASE_URL = 'http://localhost:8080/api';

// 获取平台统计数据 (拍卖总数、出价总数)
export const fetchStats = async () => {
  const res = await fetch(`${API_BASE_URL}/stats`);
  const json = await res.json();
  return json.data;
};

// 👉 升级：支持传入 status 参数 (active / ended) 进行过滤
export const fetchAuctions = async (status: string = 'active') => {
  const res = await fetch(`${API_BASE_URL}/auctions?status=${status}`);
  const json = await res.json();
  return json.data || []; // 加个兜底，防止后端返回 null 导致前端报错
};

// 👉 新增：获取某个特定拍卖的出价历史记录
export const fetchBidHistory = async (auctionId: string | number) => {
  const res = await fetch(`${API_BASE_URL}/auctions/${auctionId}/bids`);
  const json = await res.json();
  return json.data || []; // 同样加上兜底
};