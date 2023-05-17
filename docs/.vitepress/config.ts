import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base: "/naro-text/",
  title: "なろう講習会",
  description: "Webエンジニアになろう講習会のテキスト",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [{ text: "Home", link: "/" }],

    sidebar: [
      {
        text: "第一部",
        items: [{ text: "はじめに", link: "/chapter1/index" }],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/traPtitech/naro-text" },
    ],
  },
});
