import { defineConfig } from 'vitepress'
import { chapter1SidebarItems } from './sidebarConfigs/chapters/chapter1/chapter1'
import { chapter2SidebarItems } from './sidebarConfigs/chapters/chapter2/chapter2'
import { chapter4SidebarItems } from './sidebarConfigs/chapters/chapter4/chapter4'
import { webBasicSidebarItems } from './sidebarConfigs/chapters/webBasic/webBasic'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base: '/naro-text-rs/',
  title: 'なろう講習会 in Rust',
  description: 'Webエンジニアになろう講習会のテキスト in Rust',
  head: [['link', { rel: 'icon', href: '/naro-text-rs/favicon.ico' }]],
  markdown: {
    theme: {
      light: 'github-dark',
      dark: 'github-dark'
    }
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [{ text: 'Home', link: '/' }],

    sidebar: {
      '/web_basic/': webBasicSidebarItems,
      '/chapter1/': chapter1SidebarItems,
      '/chapter2/': chapter2SidebarItems,
      '/chapter4/': chapter4SidebarItems
    },
    socialLinks: [{ icon: 'github', link: 'https://github.com/traP-jp/naro-text-rs' }],
    search: {
      provider: 'local',
      options: {
        translations: {
          button: {
            buttonText: '検索',
            buttonAriaLabel: '検索ボックスを開く'
          },
          modal: {
            displayDetails: '詳細を表示',
            resetButtonTitle: 'リセット',
            backButtonTitle: '戻る',
            noResultsText: '見つかりませんでした',
            footer: {
              selectText: '選択',
              selectKeyAriaLabel: '結果を選択するには、上下キーを使用します',
              navigateText: '移動',
              navigateUpKeyAriaLabel: '前の結果に移動するには、上キーを使用します',
              navigateDownKeyAriaLabel: '次の結果に移動するには、下キーを使用します',
              closeText: '閉じる'
            }
          }
        }
      }
    }
  }
})
