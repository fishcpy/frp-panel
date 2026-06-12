import React from 'react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
  SidebarProvider,
} from "@/components/ui/sidebar"
import { Toaster } from './ui/sonner'
import { ThemeProvider } from 'next-themes'

const queryClient = new QueryClient()

export const Providers = ({ children }: { children: React.ReactNode }) => {
  return (
    <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
      <SidebarProvider>
        <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
        <Toaster />
      </SidebarProvider>
    </ThemeProvider>
  )
}
