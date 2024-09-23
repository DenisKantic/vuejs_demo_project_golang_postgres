<template>
  <div class="w-full">
    <table class="table w-full mt-10 text-md" v-if="items.length > 0">
      <thead>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Price</th>
          <th>Quantity</th>
          <th>Created</th>
          <th>Updated</th>
          <th>Actions</th>
          <!-- Added header for actions -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>{{ item.id }}</td>
          <td>{{ item.title }}</td>
          <td>{{ item.price }}</td>
          <td>{{ item.quantity }}</td>
          <td>{{ item.created_at }}</td>
          <td>{{ item.updated_at }}</td>
          <td class="flex flex-row gap-10">
            <button @click="editItem(item.id)">Edit</button>
            <button @click="deleteItem(item.id)">Delete</button>
          </td>
          <!-- Added buttons for actions -->
        </tr>
      </tbody>
    </table>
    <p v-else>No items found.</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'
import { useToast } from 'vue-toast-notification'

export default defineComponent({
  name: 'ListItems',
  setup() {
    const toast = useToast({
      position: 'top-right'
    })
    const items = ref<
      Array<{
        id: number
        title: string
        description: string
        price: number
        quantity: number
        created_at: string
        updated_at: string
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

    const editItem = (id: number) => {
      // Handle the edit action
      console.log('Edit item with ID:', id)
      // You might want to redirect to an edit page or open a modal for editing
    }

    const deleteItem = async (id: number) => {
      try {
        await axios.delete(`http://localhost:8080/deleteItem?id=${id}`) // using "ID" as query parameter
        items.value = items.value.filter((item) => item.id !== id) // Update the local state
        toast.success('List item deleted')
        console.log('Deleted item with ID:', id)
      } catch (error) {
        toast.error('Error while deleting item')
        console.error('Error deleting item:', error)
      }
    }

    onMounted(fetchItems)

    return {
      items,
      editItem,
      deleteItem
    }
  }
})
</script>
