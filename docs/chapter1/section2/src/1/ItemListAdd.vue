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
const newItemName = ref('') // [!code ++]
const newItemPrice = ref(0) // [!code ++]

const addItem = () => { // [!code ++]
  items.value.push({ name: newItemName.value, price: newItemPrice.value }) // [!code ++]
} // [!code ++]
</script>

<template>
  <div>
    <div>ItemList</div>
    <ul>
      <li v-for="item in items" :key="item.name">
        <div>名前: {{ item.name }}</div>
        <div>{{ item.price }} 円</div>
      </li>
    </ul>
    <div><!-- [!code ++] -->
      <label><!-- [!code ++] -->
        名前<!-- [!code ++] -->
        <input v-model="newItemName" type="text" /><!-- [!code ++] -->
      </label><!-- [!code ++] -->
      <label><!-- [!code ++] -->
        価格<!-- [!code ++] -->
        <input v-model="newItemPrice" type="number" /><!-- [!code ++] -->
      </label><!-- [!code ++] -->
      <button @click="addItem">追加</button><!-- [!code ++] -->
    </div><!-- [!code ++] -->
  </div>
</template>

<style></style>
