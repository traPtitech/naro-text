<!-- src/components/ItemList.vue -->
<script setup lang="ts">
import { ref } from 'vue' // [!code ++]

interface Item { // [!code ++]
  name: string // [!code ++]
  price: number // [!code ++]
} // [!code ++]

const items = ref<Item[]>([ // [!code ++]
  { name: 'たまご', price: 100 }, // [!code ++]
  { name: 'りんご', price: 160 } // [!code ++]
]) // [!code ++]
</script>

<template>
  <div>ItemList</div>
</template>

<style></style>
