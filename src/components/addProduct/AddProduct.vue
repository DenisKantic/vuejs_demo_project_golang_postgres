<template>
  <div class="h-[100svh] flex justify-center items-center w-full bg-[#ffffff] px-14">
    <form @submit.prevent="submitForm" class="w-[50%] p-5 bg-gray-200 rounded-xl min-h-[50svh]">
      <p class="py-2 text-center text-2xl">Add Product Item</p>

      <div class="mb-4">
        <label class="block text-gray-700" for="title">Title</label>
        <input v-model="form.title" type="text" id="title" class="input input-bordered w-full" />
      </div>

      <div class="mb-4">
        <label class="block text-gray-700" for="description">Description</label>
        <textarea
          v-model="form.description"
          id="description"
          class="input input-bordered w-full"
        ></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700" for="quantity">Quantity</label>
        <input
          v-model.number="form.quantity"
          type="number"
          id="quantity"
          class="input input-bordered w-full"
        />
      </div>

      <div class="mb-4">
        <label class="block text-gray-700" for="price">Price</label>
        <input
          v-model.number="form.price"
          type="number"
          id="price"
          class="input input-bordered w-full"
        />
      </div>

      <button type="submit" class="btn bg-[#6870f0] w-[50%] mx-auto text-white">Submit</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toast-notification'

export default defineComponent({
  name: 'AddProduct',
  setup() {
    const router = useRouter()
    const toast = useToast({
      position: 'top-right'
    }) // Initialize the toast

    const form = ref({
      title: '',
      description: '',
      quantity: 0,
      price: 0
    })

    const submitForm = async () => {
      //   // Validation checks
      //   if (!form.value.title) {
      //     toast.error('Title is required.')
      //     return
      //   }
      //   if (!form.value.description) {
      //     toast.error('Description is required.')
      //     return
      //   }
      //   if (form.value.quantity <= 0) {
      //     toast.error('Quantity must be greater than zero.')
      //     return
      //   }
      //   if (form.value.price <= 0) {
      //     toast.error('Price must be greater than zero.')
      //     return
      //   }
      try {
        const formData = new FormData()
        formData.append('title', form.value.title)
        formData.append('description', form.value.description)
        formData.append('quantity', form.value.quantity)
        formData.append('price', form.value.price)

        const response = await axios.post('http://localhost:8080/createListItem', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        router.push('/myProfile')
        toast.success('Item created successfully!')

        console.log('Item created:', response.data)
      } catch (error) {
        console.error('Error creating item:', error)

        // Check if the error response exists and has data
        if (error.response && error.response.data && error.response.data.error) {
          // Show the first error message from the server
          toast.error(error.response.data.error.join(', ')) // Join multiple error messages
        } else {
          toast.error('An unexpected error occurred.') // Generic error message
        }
      }
    }
    return {
      form,
      submitForm
    }
  }
})
</script>
