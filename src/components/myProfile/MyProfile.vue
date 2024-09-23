<template>
  <div class="h-[100svh] px-14 bg-[#ffffff]" id="app">
    <h1 class="text-[#6870f0] text-3xl pt-5 pb-10">Products</h1>

    <div class="w-full grid grid-cols-6 gap-10 grid-rows-1">
      <input
        type="text"
        placeholder="Search here"
        class="input col-span-4 bg-[#f1f1fb] input-primary w-full"
      />
      <h1 class="text-[#6870f0] btn col-span-2 text-3xl">Add Product</h1>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'myProfile' // Use a unique name for your component
})
</script>
