<!-- src/components/ItemList.vue -->
<script setup lang="ts">
import { ref } from 'vue'

interface Item {
  name: string
  price: number
}

const items = ref<Item[]>([
  { name: 'たまご', price: 100 },
  { name: 'りんご', price: 160 }
])
</script>

<template>
  <div>
    <div>ItemList</div>
    <ul><!-- [!code ++] -->
      <li v-for="item in items" :key="item.name"><!-- [!code ++] -->
        <div>名前: {{ item.name }}</div><!-- [!code ++] -->
        <div>{{ item.price }} 円</div><!-- [!code ++] -->
      </li><!-- [!code ++] -->
    </ul><!-- [!code ++] -->
  </div>
</template>

<style></style>
