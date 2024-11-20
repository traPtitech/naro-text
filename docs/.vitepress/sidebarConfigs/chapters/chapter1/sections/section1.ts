import type { DefaultTheme } from 'vitepress'

export const section1SidebarItems: DefaultTheme.SidebarItem[] = [
  { text: '環境構築 (windows)*', link: '/chapter1/section1/0_setup-windows' },
  { text: '環境構築 (macOS)*', link: '/chapter1/section1/1_setup-unix' },
  {
    text: 'Rust で Hello World*',
    link: '/chapter1/section1/2_hello-world'
  }
]
