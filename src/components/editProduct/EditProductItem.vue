<template>
  <div class="h-[100svh] flex justify-center items-center w-full bg-[#ffffff] px-14">
    <form @submit.prevent="submitForm" class="w-[50%] p-5 bg-gray-200 rounded-xl min-h-[50svh]">
      <RouterLink to="/myProfile" class="btn bg-[#6870f0] text-white">Go Back</RouterLink>
      <p class="py-2 text-center text-2xl">Edit Product Item</p>

      <div class="mb-4">
        <label class="block text-gray-700" for="title">Title</label>
        <input
          v-model="form.title"
          type="text"
          id="title"
          class="input input-bordered mt-3 w-full"
        />
      </div>

      <div class="mb-4">
        <label class="block text-gray-700" for="description">Description</label>
        <textarea
          v-model="form.description"
          id="description"
          class="input input-bordered w-full min-h-[20svh] py-2 resize-none mt-3"
        ></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700" for="quantity">Quantity</label>
        <input
          v-model.number="form.quantity"
          type="number"
          id="quantity"
          class="input input-bordered w-full mt-3"
        />
      </div>

      <div class="mb-4">
        <label class="block text-gray-700 mt-3" for="price">Price</label>
        <input
          v-model.number="form.price"
          type="number"
          id="price"
          class="input input-bordered w-full"
        />
      </div>

      <div class="w-full flex items-center justify-center">
        <button type="submit" class="btn bg-[#6870f0] w-[50%] mx-auto text-white">Submit</button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'vue-toast-notification'

export default defineComponent({
  name: 'EditItem',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const toast = useToast({ position: 'top-right' })

    const form = ref({
      id: 0,
      title: '',
      description: '',
      quantity: 0,
      price: 0
    })

    const fetchItem = async (id: number) => {
      try {
        const response = await axios.get(`http://localhost:8080/getOneItem?id=${id}`)
        console.log(response.data)
        form.value = response.data
      } catch (error) {
        console.error('Error fetching item:', error)
        toast.error('Error fetching item details.')
      }
    }

    const validateForm = () => {
      if (!form.value.title) {
        toast.error('Title is required.')
        return false
      }
      if (!form.value.description) {
        toast.error('Description is required.')
        return false
      }
      if (form.value.quantity <= 0) {
        toast.error('Quantity must be greater than zero.')
        return false
      }
      if (form.value.price <= 0) {
        toast.error('Price must be greater than zero.')
        return false
      }
      return true
    }

    const submitForm = async () => {
      if (!validateForm()) return

      try {
        const formData = new FormData()
        formData.append('id', form.value.id.toString()) // Include ID in the formData
        formData.append('title', form.value.title)
        formData.append('description', form.value.description)
        formData.append('quantity', form.value.quantity)
        formData.append('price', form.value.price)

        await axios.put(`http://localhost:8080/updateListItem`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        toast.success('Item updated successfully!')
        router.push('/myProfile')
      } catch (error) {
        console.error('Error updating item:', error)

        // Check if the error response exists and has data
        if (error.response && error.response.data && error.response.data.error) {
          // Show the first error message from the server
          toast.error(error.response.data.error.join(', ')) // Join multiple error messages
        } else {
          toast.error('An unexpected error occurred.') // Generic error message
        }
      }
    }

    onMounted(() => {
      const id = Number(route.params.id)
      if (id) {
        fetchItem(id)
      } else {
        toast.error('Invalid item ID')
        router.push('/myProfile')
      }
    })

    return { form, submitForm }
  }
})
</script>
