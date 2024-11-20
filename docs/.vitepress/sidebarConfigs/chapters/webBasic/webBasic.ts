import type { DefaultTheme } from 'vitepress'

export const webBasicSidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: 'Web基礎講習会',
    items: [
      { text: 'はじめに', link: '/web_basic/0_index' },
      { text: '第1回 | フロントエンド', link: '/web_basic/1_frontend' },
      { text: '第2回 | バックエンド*', link: '/web_basic/2_backend' }
    ]
  }
]
