// src/services/api.ts

const API_BASE_URL = 'http://localhost:8080/api';

// 获取平台统计数据 (拍卖总数、出价总数)
export const fetchStats = async () => {
  const res = await fetch(`${API_BASE_URL}/stats`);
  const json = await res.json();
  return json.data;
};

// 获取所有活跃的拍卖列表
export const fetchAuctions = async () => {
  const res = await fetch(`${API_BASE_URL}/auctions`);
  const json = await res.json();
  return json.data;
};