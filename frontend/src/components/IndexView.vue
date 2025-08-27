<script setup lang="ts">
import { onMounted } from 'vue'
import { useIncidentsStore } from '../store/incidents'

const store = useIncidentsStore()

onMounted(() => {
  store.fetchIncidents()
})
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">All Incidents</h1>
    <ul class="space-y-4">
      <li v-for="incident in store.incidents" :key="incident.id" class="p-4 border rounded-lg bg-white shadow">
        <h2 class="font-semibold text-lg">TITULO: {{ incident.title }}</h2>
        <p v-if="incident.description" class="text-gray-700">DESC: {{ incident.description }}</p>
        <p class="text-sm text-gray-500">Type: {{ incident.incident_type }}</p>
        <p v-if="incident.location" class="text-sm text-gray-500">Location: {{ incident.location }}</p>
        <img v-if="incident.image" :src="incident.image" alt="Incident" class="mt-2 max-h-48 object-cover rounded" />
      </li>
    </ul>
  </div>
</template>