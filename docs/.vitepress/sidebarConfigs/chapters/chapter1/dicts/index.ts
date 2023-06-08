import type { DefaultTheme } from "vitepress"
import { cleanCodeSidebarItems } from "./clean-code"

export const dictSidebarItems: DefaultTheme.SidebarItem[] = [...cleanCodeSidebarItems]
