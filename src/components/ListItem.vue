<template>
  <div class="w-full">
    <div class="grid gap-5 grid-rows-1 grid-cols-6">
      <input
        type="text"
        v-model="searchQuery.title"
        placeholder="Search by title"
        class="input col-span-3 bg-[#f1f1fb] input-primary w-full"
        @input="fetchItems"
      />
      <input
        type="number"
        v-model="searchQuery.id"
        placeholder="Search by ID"
        class="input col-span-3 bg-[#f1f1fb] input-primary w-full"
        @input="fetchItems"
      />
    </div>
    <table class="table table-fixed w-full mt-10 text-md" v-if="items.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Price</th>
          <th>Quantity</th>
          <th>Description</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>{{ item.id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.price }}</td>
          <td>{{ item.quantity }}</td>
          <td>{{ formatDescription(item.description) }}</td>
          <td>{{ formatDate(item.created_at) }}</td>
        </tr>
      </tbody>
    </table>
    <p v-if="noItemsFound" class="text-red-500">No items found.</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'
import { formatDescription } from '../helper/formatDescription'
import { formatDate } from '../helper/formatDate'

export default defineComponent({
  name: 'ListItems',
  setup() {
    const items = ref([])
    const searchQuery = ref({ title: '', id: null }) // Changed to an object
    const noItemsFound = ref(false)

    const fetchItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/getListItems', {
          params: {
            title: searchQuery.value.title,
            id: searchQuery.value.id // Include ID in the params
          }
        })

        items.value = response.data.items
        noItemsFound.value =
          items.value.length === 0 &&
          (searchQuery.value.title.trim() !== '' || searchQuery.value.id !== null)

        console.log(response.data.items)
      } catch (error) {
        console.error('Error fetching items:', error)
      }
    }

    onMounted(fetchItems)

    return {
      items,
      searchQuery,
      fetchItems,
      noItemsFound,
      formatDate,
      formatDescription
    }
  }
})
</script>
