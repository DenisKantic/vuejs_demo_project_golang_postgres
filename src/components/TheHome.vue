<template>
  <main class="w-full h-screen bg-[#ffffff] px-14">
    <h1 class="text-[#6870f0] text-3xl pt-5 pb-10 cursor-default">Products</h1>
    <ListItem />
  </main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ListItem from './ListItem.vue' // Importing the ListItem component

export default defineComponent({
  name: 'TheHome', // Use a unique name for your component
  components: {
    ListItem // Declare the ListItem component
  }
})
</script>
