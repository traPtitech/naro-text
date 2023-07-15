<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/">Home</router-link>
        |
        <router-link to="/ping">Ping</router-link>
        |
        <router-link to="/login">Login</router-link>
      </header>

      <router-view />
    </div>
  </main>
</template>
