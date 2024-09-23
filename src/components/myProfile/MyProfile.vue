<template>
  <div class="h-[100svh] px-14 bg-[#ffffff]">
    <h1 class="text-[#6870f0] text-3xl text-center pt-5 pb-10 cursor-default">Dashboard</h1>

    <h1 class="text-[#6870f0] text-3xl pt-5 pb-10 cursor-default">Products</h1>

    <div class="w-full grid grid-cols-6 gap-10 grid-rows-1">
      <input
        type="text"
        placeholder="Search here"
        class="input col-span-5 bg-[#f1f1fb] input-primary w-full"
      />
      <RouterLink to="/addProduct" class="btn bg-[#6870f0] text-white col-span-1 text-xl">
        Add Product <i class="fa-solid fa-plus text-xl"></i>
      </RouterLink>
    </div>
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
