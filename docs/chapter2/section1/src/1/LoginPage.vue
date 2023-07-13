<script setup>
import axios from "axios";
import { ref } from 'vue';

const username = ref('');
const password = ref('');
const login = () => axios.post("/api/login", { username: username.value, password: password.value });

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