"use client"

import { ChevronRight, type LucideIcon } from "lucide-react"
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar"

export function NavMain({
  items,
}: {
  items: {
    title: string
    url: string
    icon?: LucideIcon
    isActive?: boolean
    items?: {
      title: string
      url: string
    }[]
  }[]
}) {
  const router = useRouter()
  const [openGroups, setOpenGroups] = useState<Record<string, boolean>>({})

  const urlSelected = (url: string) => {
    if (typeof window !== "undefined") {
      const pathname = window.location.pathname
      return pathname === url
    }
  }

  const isGroupActive = (item: typeof items[0]) => {
    if (!item.items) return false
    if (typeof window !== "undefined") {
      const pathname = window.location.pathname
      // 检查当前路径是否匹配任何子菜单项
      return item.items.some(subItem => pathname === subItem.url || pathname.startsWith(subItem.url + '/'))
    }
    return false
  }

  // 初始化和路由变化时更新展开状态
  useEffect(() => {
    const newOpenGroups: Record<string, boolean> = {}
    items.forEach(item => {
      if (item.items) {
        newOpenGroups[item.title] = isGroupActive(item)
      }
    })
    setOpenGroups(newOpenGroups)
  }, [router.pathname])

  return (
    <SidebarMenu>
      {items.map((item) => (
        <Collapsible
          key={item.title}
          asChild
          open={openGroups[item.title]}
          onOpenChange={(open) => setOpenGroups(prev => ({ ...prev, [item.title]: open }))}
          className="group/collapsible"
        >
          <>
            {!item.items && <SidebarMenuItem>
              <SidebarMenuButton isActive={urlSelected(item.url)} onClick={() => router.push(item.url)} tooltip={item.title}>
                {item.icon && <item.icon />}
                <span>{item.title}</span>
              </SidebarMenuButton>
            </SidebarMenuItem>
            }
            {item.items && <SidebarMenuItem>
              <CollapsibleTrigger asChild>
                <SidebarMenuButton tooltip={item.title}>
                  {item.icon && <item.icon />}
                  <span>{item.title}</span>
                  <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
                </SidebarMenuButton>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <SidebarMenuSub>
                  {item.items?.map((subItem) => (
                    <SidebarMenuSubItem key={subItem.title}>
                      <SidebarMenuSubButton
                        isActive={urlSelected(subItem.url)}
                        onClick={() => router.push(subItem.url)}
                      >
                        <span>{subItem.title}</span>
                      </SidebarMenuSubButton>
                    </SidebarMenuSubItem>
                  ))}
                </SidebarMenuSub>
              </CollapsibleContent>
            </SidebarMenuItem>}
          </>
        </Collapsible>
      ))}
    </SidebarMenu>
  )
}
