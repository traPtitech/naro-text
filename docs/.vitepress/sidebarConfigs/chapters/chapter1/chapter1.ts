import type { DefaultTheme } from 'vitepress'
import { section1SidebarItems } from './sections/section1'
import { section2SidebarItems } from './sections/section2'
import { section3SidebarItems } from './sections/section3'
import { section4SidebarItems } from './sections/section4'
import { dictSidebarItems } from './dicts'

export const chapter1SidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: '第一部',
    items: [
      { text: 'はじめに', link: '/chapter1/index' },
      ...dictSidebarItems,
      {
        text: '実習編',
        collapsed: false,
        items: [
          ...section1SidebarItems,
          ...section2SidebarItems,
          ...section3SidebarItems,
          ...section4SidebarItems,
          {
            text: 'React入門',
            link: '/chapter1/dicts/react/0_react-intro'
          },
          {
            text: 'アプリを作ってみよう(React)',
            link: '/chapter1/dicts/react/1_create-app'
          },
          {
            text: '制作物を公開してみよう(React)',
            link: '/chapter1/dicts/react/2_publish'
          }
        ]
      }
    ]
  }
]
