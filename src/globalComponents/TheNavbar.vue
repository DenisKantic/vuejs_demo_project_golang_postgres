<template>
  <div class="navbar fixed bg-gray-300 shadow-md px-14">
    <div class="flex-1">
      <RouterLink to="/" class="text-xl text-black font-bold">Home</RouterLink>
    </div>
    <div class="flex-none gap-2">
      <div class="dropdown dropdown-end">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
          <div class="w-10 rounded-full">
            <img
              alt="Tailwind CSS Navbar component"
              src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
            />
          </div>
        </div>
        <ul
          tabindex="0"
          class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
        >
          <li>
            <RouterLink to="myProfile" class="justify-between"> My Profile </RouterLink>
          </li>
          <li><a>Add Product</a></li>
          <li><a>Logout</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TheNavbar' // Adjust the name as necessary
})
</script>
