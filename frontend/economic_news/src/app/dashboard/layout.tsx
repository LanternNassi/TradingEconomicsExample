import Navbar from "@/components/Navbar";


export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
      
        <div className="flex flex-col min-h-screen">
            <Navbar/>
            <main className="flex-grow container mx-auto p-4">
                {children}
            </main>
        </div>
  );
}
