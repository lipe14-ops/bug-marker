<template>
  <div class="flex flex-col min-h-0">
    <div class="flex-shrink-0 p-4 border-b border-white/10">
      <div class="flex justify-between items-center mb-4">
        <div class="space-y-1">
          <h3 class="text-white font-bold text-lg">Polígonos</h3>
          <p class="text-gray-400 text-xs">{{ polygons.length }} polígonos marcados</p>
        </div>
        <button 
          class="flex items-center justify-center gap-2 rounded-lg h-9 px-4 bg-primary text-background-dark text-sm font-bold hover:bg-primary/90 transition-colors"
          @click="$emit('new', imageID)"
        >
          <span class="material-symbols-outlined !text-xl">add</span>
          <span>Novo Polígono</span>
        </button>
      </div>
    </div>
    <div class="flex-1 min-h-0 overflow-y-auto px-4 pb-4 space-y-4 custom-scrollbar">
      <BmPolygonItem
        v-for="polygon in polygons"
        :name="polygon.name"
        :id="polygon.id"
        :class-name="polygon.className"
        :class-color="classes[polygon.classID]?.color || '#00000000'"
        :points="polygon.coordinates"
        :is-open="polygon.id === openPolygonId"
        :is-editing="polygon.id === editingPolygonId"
        @edit="$emit('edit', polygon)"
        @delete="$emit('delete', polygon.id)"
        @add-point="$emit('add-point', polygon)"
      
      />
    </div>
  </div>    
</template>

<script setup>
import { ref } from 'vue'
import BmPolygonItem from '../atoms/BmPolygonItem.vue'
import { useTokenStore } from '../../store.js'
import BmPolygonActionButton from '../atoms/BmPolygonActionButton.vue'
const tokenStore = useTokenStore()


const props = defineProps({
  currentImage: {
    type: String,
    required: true
  },
  imageID: {
    type: String,
    required: true
  },
  projectID: {
    type: String,
    required: true
  },
  polygons: {
    type: Array,
    required: true,
    default: () => []
  },
  openPolygonId: {
    type: Number,
    default: null
  }
  ,
  editingPolygonId: {
    type: [String, Number],
    default: null
  }
})

const classes = ref({})

  fetch(`http://localhost:8000/api/project/${props.projectID}/classes`, {
    method: "GET", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON      
  .then(data => {
    for (let classe of data.data) {
      classes.value[classe.id] = classe
    }
    console.log(classes.value)
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  })


console.log('Polygons prop:', props.polygons)

defineEmits(['new', 'edit', 'delete', 'add-point'])
</script>