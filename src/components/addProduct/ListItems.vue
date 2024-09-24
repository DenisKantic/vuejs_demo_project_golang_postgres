<template>
  <div class="w-full">
    <div class="mb-4">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search by title"
        class="input input-bordered w-full"
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
          <td class="flex flex-row gap-10">
            <button @click="editItem(item.id)" class="btn btn-neutral">Edit</button>
            <button @click="deleteItem(item.id)" class="btn btn-error text-white">Delete</button>
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
import { useToast } from 'vue-toast-notification'
import { useRouter } from 'vue-router'
import { formatDate } from '../../helper/formatDate'
import { formatDescription } from '../../helper/formatDescription'

export default defineComponent({
  name: 'ListItems',
  setup() {
    const toast = useToast({
      position: 'top-right'
    })
    const router = useRouter()
    const items = ref<
      Array<{
        id: number
        title: string
        description: string
        price: number
        quantity: number
        created_at: string
      }>
    >([])
    const searchQuery = ref('') // New reactive property for search

    const fetchItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/getListItems', {
          params: { title: searchQuery.value } // Pass the search query as a parameter
        })
        items.value = response.data.items
        console.log(response.data.items)
      } catch (error) {
        console.error('Error fetching items:', error)
      }
    }

    const editItem = (id: number) => {
      // Handle the edit action
      console.log('Edit item with ID:', id)
      try {
        router.push({ name: 'editItem', params: { id } })
      } catch (error) {
        toast.error('Error for editing item')
        console.log(error)
      }
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
      searchQuery,
      fetchItems, // Return fetchItems to be callable on input event
      editItem,
      deleteItem,
      formatDate,
      formatDescription
    }
  }
})
</script>
