<script setup lang="ts">
import { ref } from "vue"

interface Item {
  name: string
  price: number
}

const items = ref<Item[]>([
  { name: "たまご", price: 100 },
  { name: "りんご", price: 160 }
])
</script>

<template>
  <div>ItemList</div>
</template>

<style></style>
