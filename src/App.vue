<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import '@fortawesome/fontawesome-free/css/all.css'
import Navbar from './globalComponents/TheNavbar.vue'
</script>

<template>
  <div class="text-xl h-screen w-full bg-purple-100">
    <Navbar />
    <RouterView class="pt-20" />
  </div>
</template>

<style scoped></style>
