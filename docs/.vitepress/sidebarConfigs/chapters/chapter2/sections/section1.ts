import type { DefaultTheme } from 'vitepress'

// export const section1SidebarItems: DefaultTheme.SidebarItem[] = [
//   { text: 'サーバーのログイン機能の実装', link: '/chapter2/section1/0_server_setup' },
//   { text: 'vue-routerとプロキシの設定', link: '/chapter2/section1/1_router-setup' },
//   { text: 'エンドポイントにアクセスする', link: '/chapter2/section1/2_axios' }
// ]

export const section1SidebarItems: DefaultTheme.SidebarItem[] = [
  { text: 'プロジェクトのセットアップ*', link: '/chapter2/section1/0_setup' },
  { text: 'アカウント機能の実装*', link: '/chapter2/section1/1_account' },
  { text: 'セッションの実装*', link: '/chapter2/section1/2_session' },
  { text: '検証*', link: '/chapter2/section1/3_verify' },
  { text: 'おまけ演習問題', link: '/chapter2/section1/4_extra' }
]
