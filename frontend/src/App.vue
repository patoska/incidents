<template>
  <div class="min-h-screen bg-gray-100 text-gray-900">
    <nav class="bg-indigo-600 text-white p-4 flex gap-4">
      <RouterLink to="/" class="hover:underline">Incidents</RouterLink>
      <RouterLink to="/new" class="hover:underline">New Incident</RouterLink>
    </nav>
    <main class="p-6">
      <RouterView />
    </main>
  </div>
</template>