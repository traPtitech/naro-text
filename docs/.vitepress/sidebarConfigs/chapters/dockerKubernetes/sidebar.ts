import type { DefaultTheme } from 'vitepress'

export const dockerKubernetesSidebarItems: DefaultTheme.SidebarItem[] = [
  {
    text: 'Docker・Kubernetes講習会',
    items: [
      { text: 'はじめに', link: '/docker_kubernetes/0_index' },
      { text: '事前準備', link: '/docker_kubernetes/1_prepare' },
      { text: '第1回 | Docker・Compose', link: '/docker_kubernetes/2_docker_compose' },
      { text: '第2回 | Kubernetesの基本（1）', link: '/docker_kubernetes/3_kubernetes_basics' },
      { text: '第3回 | Kubernetesの基本（2）', link: '/docker_kubernetes/4_configuration' },
      { text: '講習後に学ぶこと', link: '/docker_kubernetes/5_extra' }
    ]
  }
]
