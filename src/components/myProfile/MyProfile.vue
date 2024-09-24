<template>
  <div class="h-[100svh] px-14 bg-[#ffffff]">
    <h1 class="text-[#6870f0] text-3xl text-center pt-5 pb-10 cursor-default">Dashboard</h1>

    <h1 class="text-[#6870f0] text-3xl pt-5 pb-10 cursor-default">Products</h1>

    <ListComponent />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ListComponent from '../addProduct/ListItems.vue'

export default defineComponent({
  name: 'myProfile', // Use a unique name for your component
  components: {
    ListComponent // Declare the ListItem component
  }
})
</script>
