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
        class="input col-span-2 bg-[#f1f1fb] input-primary w-full"
        @input="fetchItems"
      />
      <RouterLink to="/addProduct" class="btn bg-[#6870f0] col-span-1 text-white text-xl">
        Add Product <i class="fa-solid fa-plus text-xl"></i>
      </RouterLink>
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
    <p v-if="noItemsFound" class="text-red-500">No items found.</p>
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

    const editItem = (id: number) => {
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
        await axios.delete(`http://localhost:8080/deleteItem?id=${id}`)
        items.value = items.value.filter((item) => item.id !== id)
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
      fetchItems,
      editItem,
      noItemsFound,
      deleteItem,
      formatDate,
      formatDescription
    }
  }
})
</script>
