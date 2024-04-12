import type { DefaultTheme } from 'vitepress'
import { sidebarItems } from './sidebar'

export const chapter4SidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: '第四部',
    items: [{ text: 'はじめに', link: '/chapter4/0_index' }, ...sidebarItems]
  }
]
