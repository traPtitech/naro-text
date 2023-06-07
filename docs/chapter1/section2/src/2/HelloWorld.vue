<script setup lang="ts">
import ClickCounter from './ClickCounter.vue'
import ItemList from './ItemList.vue'

defineProps<{
  msg: string
}>()
</script>

<template>
  <div>
    <h1>{{ msg }}</h1>
    <ClickCounter />
    <ItemList />
  </div>
</template>

<style scoped>
a {
  color: #42b983;
}
</style>
