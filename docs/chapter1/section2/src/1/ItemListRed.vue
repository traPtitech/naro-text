<template>
  <div>
    <div>ItemList</div>
    <ul>
      <li v-for="item in items" :key="item.name" :class="{ over500: item.price >= 500 }">
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
      <button @click="addItem">add</button>
    </div>
  </div>
</template>

<style>
.over500 {
  color: red;
}
</style>
