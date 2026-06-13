"use client"

import { Toaster } from "@/components/ui/sonner"
import { AppSidebar } from "@/components/app-sidebar"
import { Separator } from "@/components/ui/separator"
import {
  SidebarInset,
  SidebarTrigger,
} from "@/components/ui/sidebar"

export function RootLayout({
  children,
  sidebarItems,
  sidebarFooter,
  mainHeader,
}: {
  children: React.ReactNode;
  sidebarItems?: React.ReactNode;
  sidebarFooter?: React.ReactNode;
  mainHeader?: React.ReactNode
}) {
  return (
    <>
      <AppSidebar footer={sidebarFooter}>{sidebarItems}</AppSidebar>
      <SidebarInset>
        <header className="flex h-12 shrink-0 items-center gap-2 border-b px-4">
          <SidebarTrigger className="-ml-1" />
          <Separator orientation="vertical" className="mr-2 h-4" />
          {mainHeader}
        </header>
        <div className="flex flex-1 flex-col gap-4 p-4 overflow-auto">
          {children}
        </div>
        <Toaster />
      </SidebarInset>
    </>
  )
}