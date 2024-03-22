<script setup lang="ts">
import { ref } from 'vue'

const username = ref<string>('')
const password = ref<string>('')
const login = () =>
  fetch('/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ username: username.value, password: password.value })
  })
</script>

<template>
  <div class="login">
    <h1>This is an login page</h1>
    <div>
      <input type="text" v-model="username" />
      <input type="password" v-model="password" />
    </div>
    <div>
      <button @click="login">login</button>
    </div>
  </div>
</template>
