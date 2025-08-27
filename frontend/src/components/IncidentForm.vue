<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useIncidentsStore, type Incident } from '../store/incidents'

const store = useIncidentsStore()
const router = useRouter()

const form = ref<Incident>({
  title: '',
  description: '',
  incident_type: 'fire',
  location: '',
  image: '',
})

const incidentTypes = ['fire', 'accident', 'theft']

function handleFileUpload(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) return

  const reader = new FileReader()
  reader.onload = () => {
    form.value.image = reader.result as string
  }
  reader.readAsDataURL(file)
}

async function submitForm() {
  if (!form.value.title || !form.value.incident_type) {
    alert('Title and Incident type are required.')
    return
  }

  await store.createIncident(form.value)
  router.push('/')
}
</script>

<template>
  <form @submit.prevent="submitForm" class="space-y-4 max-w-lg">
    <div>
      <label class="block mb-1 font-medium">Title *</label>
      <input v-model="form.title" class="w-full border rounded p-2" required />
    </div>

    <div>
      <label class="block mb-1 font-medium">Description</label>
      <textarea v-model="form.description" class="w-full border rounded p-2" />
    </div>

    <div>
      <label class="block mb-1 font-medium">Incident Type *</label>
      <select v-model="form.incident_type" class="w-full border rounded p-2">
        <option v-for="type in incidentTypes" :key="type" :value="type">{{ type }}</option>
      </select>
    </div>

    <div>
      <label class="block mb-1 font-medium">Location</label>
      <input v-model="form.location" class="w-full border rounded p-2" />
    </div>

    <div>
      <label class="block mb-1 font-medium">Image</label>
      <input type="file" @change="handleFileUpload" />
      <img v-if="form.image" :src="form.image" class="mt-2 max-h-32 rounded shadow" />
    </div>

    <button type="submit" class="px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700">
      Submit
    </button>
  </form>
</template>
