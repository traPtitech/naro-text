<script setup lang="ts">
import axios from 'axios'
const url = 'https://ent5o104rfsmh.x.pipedream.net'
const sendGet = () => axios.get(url)
const sendPost = () =>
  axios.post(url, {
    traQID: 'noc7t',
    grade: 22
  })
const sendPut = () =>
  axios.put(url, {
    testData: { string: 'string', number: 1 }
  })
const sendDelete = () => axios.delete(url, {})
</script>

<template>
  <div>
    <button @click="sendGet">GET</button>
    <button @click="sendPost">POST</button>
    <button @click="sendPut">PUT</button>
    <button @click="sendDelete">DELETE</button>
  </div>
</template>
