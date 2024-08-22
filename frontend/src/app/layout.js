
import "./globals.css";
import { Sidebar } from '@/components/Sidebar';
import { Footer } from '@/components/Footer';
import { Toolbar } from "@/components/Toolbar";


export const metadata = {
  title: "Kubestatus",
  description: "Kubernetes Looking Glass",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className="relative bg-gray-50 overflow-hidden max-h-screen">
        <toolbar id="navbar">
          <Toolbar />
        </toolbar>

        <div>

          <aside id="sidebar" className="fixed bg-red-200 shadow-md max-h-screen w-60">
            <Sidebar />
          </aside>

          {children}
 
        </div>
        
        <footer id="footer">
          <Footer />
        </footer>

      </body>
    </html>
  );
}
