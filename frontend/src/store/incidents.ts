import { defineStore } from 'pinia'

export interface Incident {
  id?: number
  title: string
  description?: string
  incident_type: string
  location?: string
  image?: string
}

export const useIncidentsStore = defineStore('incidents', {
  state: () => ({
    incidents: [] as Incident[],
  }),

  actions: {
    async fetchIncidents() {
      const res = await fetch(`${import.meta.env.VITE_API_ENDPONT}/api/incidents`)
      console.log("res")
      console.log(res)
      this.incidents = await res.json()
      console.log("this.incidents")
      console.log(this.incidents)
    },

    async createIncident(incident: Incident) {
      const res = await fetch(`${import.meta.env.VITE_API_ENDPONT}/api/incidents`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(incident),
      })
      const newIncident = await res.json()
      // console.log("newIncident")
      // console.log(newIncident)
      if (!Array.isArray(this.incidents)) {
        this.incidents = []
      }
      this.incidents.push(newIncident)
    },
  },
})