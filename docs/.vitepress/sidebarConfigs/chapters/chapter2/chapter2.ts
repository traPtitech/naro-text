import type { DefaultTheme } from 'vitepress'
import { section1SidebarItems } from './sections/section1'
import { section2SidebarItems } from './sections/section2'
import { section3SidebarItems } from './sections/section3'
import { section4SidebarItems } from './sections/section4'
import { dictSidebarItems } from './dicts'

export const chapter2SidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: '第二部',
    items: [
      { text: 'はじめに', link: '/chapter2/index' },
      ...dictSidebarItems,
      {
        text: '実習編',
        collapsed: false,
        items: [
          ...section1SidebarItems,
          ...section2SidebarItems,
          ...section3SidebarItems,
          ...section4SidebarItems
        ]
      }
    ]
  }
]
