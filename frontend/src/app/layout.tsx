import "./globals.css";
import { Web3Provider } from "@/components/Web3Provider"; 
// 👉 1. 引入 react-hot-toast 的全局容器
import { Toaster } from "react-hot-toast";
// 👉 2. 引入我们刚刚新建的网络检查组件 (请确保路径与你实际建的一致)
import NetworkChecker from "../components/NetworkChecker"; 

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="zh">
      <body>
        {/* 👉 全局的轻提示容器，放在最外层即可 */}
        <Toaster position="top-right" reverseOrder={false} />
        
        <Web3Provider>
          {/* 👉 网络检查器必须放在 Web3Provider 内部 */}
          <NetworkChecker />
          {children}
        </Web3Provider>
      </body>
    </html>
  );
}