import type { DefaultTheme } from "vitepress"

export const section2SidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: "Vue 入門",
    link: "/chapter1/section2/index",
    items: [
      {
        text: "Vueに触れてみる",
        link: "/chapter1/section2/1"
      },
      {
        text: "もっと触れてみる",
        link: "/chapter1/section2/2"
      },
      {
        text: "公開してみよう",
        link: "/chapter1/section2/3"
      }
    ]
  }
]
