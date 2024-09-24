<template>
  <div class="w-full">
    <table class="table table-fixed w-full mt-10" v-if="items.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Price</th>
          <th>Quantity</th>
          <th>Description</th>
          <th>Created</th>
          <th>Actions</th>
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
          <td>
            <button>Read More</button>
          </td>
        </tr>
      </tbody>
    </table>
    <p v-else>No items found.</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'
import { formatDate } from '../helper/formatDate'
import { formatDescription } from '../helper/formatDescription'

export default defineComponent({
  name: 'ListItems',
  setup() {
    const items = ref<
      Array<{
        id: number
        title: string
        description: string
        price: number
        quantity: number
        description: string
        created_at: string
      }>
    >([])

    const fetchItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/getListItems') // Adjust the URL as needed
        items.value = response.data.items
        console.log(response.data.items) // Assuming your API response directly contains the item array
      } catch (error) {
        console.error('Error fetching items:', error)
      }
    }

    onMounted(fetchItems)

    return {
      items,
      formatDate, // make it visible in template
      formatDescription
    }
  }
})
</script>
