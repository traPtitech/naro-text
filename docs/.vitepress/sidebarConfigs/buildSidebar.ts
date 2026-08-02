import type { DefaultTheme } from 'vitepress'
import { chapter1SidebarItems } from './chapters/chapter1/chapter1'
import { chapter2SidebarItems } from './chapters/chapter2/chapter2'
import { chapter4SidebarItems } from './chapters/chapter4/chapter4'
import { webBasicSidebarItems } from './chapters/webBasic/webBasic'

type ChapterKey = 'web_basic' | 'chapter1' | 'chapter2' | 'chapter4'

const chapterLinks: Record<ChapterKey, DefaultTheme.SidebarItem> = {
  web_basic: { text: 'Web基礎講習会', link: '/web_basic/0_index' },
  chapter1: { text: '第一部', link: '/chapter1/index' },
  chapter2: { text: '第二部', link: '/chapter2/index' },
  chapter4: { text: '第四部', link: '/chapter4/0_index' }
}

const chapterFullItems: Record<ChapterKey, DefaultTheme.SidebarItem[]> = {
  web_basic: webBasicSidebarItems,
  chapter1: chapter1SidebarItems,
  chapter2: chapter2SidebarItems,
  chapter4: chapter4SidebarItems
}

const chapterOrder: ChapterKey[] = ['web_basic', 'chapter1', 'chapter2', 'chapter4']

export function buildSidebar(current: ChapterKey): DefaultTheme.SidebarItem[] {
  return [
    { text: 'トップページ', link: '/' },
    { text: '座学編（スライド）', link: '/lectures/' },
    ...chapterOrder.flatMap((chapter) =>
      chapter === current ? chapterFullItems[chapter] : [chapterLinks[chapter]]
    )
  ]
}
