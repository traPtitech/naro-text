<!-- src/components/ItemList.vue -->
<script setup lang="ts">
import { ref } from 'vue'

interface Item {
  name: string
  price: number
}

const items = ref<Item[]>([
  { name: 'たまご', price: 100 },
  { name: 'りんご', price: 160 },
])
const newItemName = ref('')
const newItemPrice = ref(0)

const addItem = () => {
  if (items.value.some((item) => item.name === newItemName.value)) return // [!code ++]
  items.value.push({ name: newItemName.value, price: newItemPrice.value })
}

const deleteItem = (name: string) => {
  items.value = items.value.filter((item) => item.name !== name)
}
</script>

<template>
  <div>
    <div>ItemList</div>
    <ul>
      <li v-for="item in items" :key="item.name">
        <div>名前: {{ item.name }}</div>
        <div>{{ item.price }} 円</div>
        <button @click="deleteItem(item.name)">削除</button>
      </li>
    </ul>
    <div>
      <label>
        名前
        <input v-model="newItemName" type="text" />
      </label>
      <label>
        価格
        <input v-model="newItemPrice" type="number" />
      </label>
      <button @click="addItem">追加</button>
    </div>
  </div>
</template>

<style></style>
