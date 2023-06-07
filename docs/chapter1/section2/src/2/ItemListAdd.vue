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
const newItemName = ref("")
const newItemPrice = ref(0)

const addItem = () => {
  items.value.push({ name: newItemName.value, price: newItemPrice.value })
}
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
