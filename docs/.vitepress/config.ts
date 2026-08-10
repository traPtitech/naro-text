import { defineConfig } from 'vitepress'
import { buildSidebar } from './sidebarConfigs/buildSidebar'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base: '/naro-text/',
  title: 'なろう講習会',
  description: 'Webエンジニアになろう講習会のテキスト',
  head: [
    ['link', { rel: 'icon', href: '/naro-text/favicon.ico' }],
    ['link', { rel: 'preconnect', href: 'https://fonts.googleapis.com' }],
    ['link', { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' }],
    [
      'link',
      {
        rel: 'stylesheet',
        href: 'https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@400;500;700&family=Noto+Sans+Mono:wght@400;600&display=swap'
      }
    ]
  ],
  markdown: {
    theme: {
      light: 'github-dark',
      dark: 'github-dark'
    }
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    siteTitle: 'なろう講習会',
    nav: [
      { text: 'Home', link: '/' },
      { text: '座学編', link: '/lectures/' }
    ],

    sidebar: {
      '/web_basic/': buildSidebar('web_basic'),
      '/chapter1/': buildSidebar('chapter1'),
      '/chapter2/': buildSidebar('chapter2'),
      '/chapter4/': buildSidebar('chapter4')
    },
    socialLinks: [{ icon: 'github', link: 'https://github.com/traPtitech/naro-text' }],
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
