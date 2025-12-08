<script setup>
import { ref } from 'vue'
import { useTokenStore } from '../../store.js'

const classes = ref([])

const props = defineProps({
  projectID: String
})

const selected = defineModel()

const tokenStore = useTokenStore()

const open = ref(false)

fetch(`http://localhost:8000/api/project/${props.projectID}/classes/`, {
  method: "GET",
  headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
  },
})
.then(response => response.json())
.then(data => {
  classes.value = data.data
  selected.value = classes.value[0]
})
.catch(error => {
  console.error('Error fetching classes:', error)
})

function selectColor(option) {
  selected.value = option
  open.value = false
  console.log('Selected class:', option)
}
</script>

<template>
  <div v-if="selected" class="relative inline-block text-left w-full">
    <!-- Botão do select -->
    <button
      @click="open = !open"
      class="flex w-full items-center justify-between w-48 bg-background-dark text-white px-4 py-2 rounded-md border border-gray-700 hover:border-gray-500 transition"
    >
      <div class="flex items-center gap-3">
        <div class="w-4 h-4 rounded-full" :style="{ backgroundColor: selected.color }"></div>
        <span class="font-medium">{{ selected.name }}</span>
      </div>
      <svg
        class="w-4 h-4 ml-2 opacity-70"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        viewBox="0 0 24 24"
      >
        <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
      </svg>
    </button>

    <!-- Opções -->
    <div
      v-if="open"
      class="absolute w-full z-10 mt-2 bg-background-dark border border-gray-700 rounded-md shadow-lg overflow-hidden"
    >
      <div
        v-for="option in classes"
        :key="option.name"
        @click="selectColor(option)"
        class="flex items-center gap-3 px-4 py-2 cursor-pointer hover:bg-gray-800 transition"
      >
        <div class="w-4 h-4 rounded-full" :style="{ backgroundColor: option.color }"></div>
        <span class="text-white font-medium">{{ option.name }}</span>
      </div>
    </div>
  </div>
</template>
