import type { DefaultTheme } from 'vitepress'
import { cleanCodeSidebarItems } from './clean-code'
import { commitSizeSidebarItems } from './commit-size'
import { DevToolsSidebarItems } from './DevTools'

export const dictSidebarItems: DefaultTheme.SidebarItem[] = [
  ...cleanCodeSidebarItems,
  ...commitSizeSidebarItems,
  ...DevToolsSidebarItems
]
