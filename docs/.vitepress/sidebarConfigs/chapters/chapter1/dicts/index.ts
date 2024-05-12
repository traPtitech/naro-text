import type { DefaultTheme } from 'vitepress'
import { cleanCodeSidebarItems } from './clean-code'
import { DevToolsSidebarItems } from './DevTools'

export const dictSidebarItems: DefaultTheme.SidebarItem[] = [
  ...cleanCodeSidebarItems,
  ...DevToolsSidebarItems
]
