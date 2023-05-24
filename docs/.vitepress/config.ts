import { defineConfig } from "vitepress"

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base: "/naro-text/",
  title: "なろう講習会",
  description: "Webエンジニアになろう講習会のテキスト",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [{ text: "Home", link: "/" }],

    sidebar: {
      "/chapter1/": [
        {
          text: "第一部",
          items: [
            { text: "はじめに", link: "/chapter1/index" },
            {
              text: "実習編",
              collapsed: true,
              items: [
                { text: "環境構築", link: "/chapter1/section1/setup" },
                {
                  text: "Golang で Hello World",
                  link: "/chapter1/section1/hello-world"
                }
              ]
            }
          ]
        }
      ]
    },
    socialLinks: [{ icon: "github", link: "https://github.com/traPtitech/naro-text" }],
    search: {
      provider: "local",
      options: {
        translations: {
          button: {
            buttonText: "検索",
            buttonAriaLabel: "検索ボックスを開く"
          },
          modal: {
            displayDetails: "詳細を表示",
            resetButtonTitle: "リセット",
            backButtonTitle: "戻る",
            noResultsText: "見つかりませんでした",
            footer: {
              selectText: "選択",
              selectKeyAriaLabel: "結果を選択するには、上下キーを使用します",
              navigateText: "移動",
              navigateUpKeyAriaLabel: "前の結果に移動するには、上キーを使用します",
              navigateDownKeyAriaLabel: "次の結果に移動するには、下キーを使用します",
              closeText: "閉じる"
            }
          }
        }
      }
    }
  }
})
