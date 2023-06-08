<script setup lang="ts">
import { ref, computed } from 'vue'

const count = ref<number>(0)
const countMessage = computed(() => '回数: ' + count.value)
</script>

<template>
  <div>
    <div>{{ countMessage }}</div>
    <button @click="count++">クリック！</button>
    <button @click="count = 0">リセット！</button>
  </div>
</template>

<style></style>
