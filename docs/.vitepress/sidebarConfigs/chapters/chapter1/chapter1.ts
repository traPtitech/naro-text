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
        text: '第3回 | 環境構築',
        collapsed: false,
        items: section1SidebarItems
      },
      {
        text: '第4回 | フロントエンド',
        collapsed: false,
        items: section2SidebarItems
      },
      {
        text: '第5回 | サーバー',
        collapsed: false,
        items: section3SidebarItems
      },
      {
        text: '第6回 | データベース',
        collapsed: false,
        items: section4SidebarItems
      },
      {
        text: 'React入門',
        link: '/chapter1/dicts/react/0_react-intro'
      },
      {
        text: 'アプリを作ってみよう(React)',
        link: '/chapter1/dicts/react/1_create-app'
      }
    ]
  }
]
