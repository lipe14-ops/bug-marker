<style>
  .material-symbols-outlined {
      font-variation-settings: 'FILL' 0, 'wght' 400, 'GRAD' 0, 'opsz' 24;
      font-size: 24px;
  }
</style>

<script setup>
import { ref } from 'vue'
import BmSidebarMenu from '../components/molecules/BmSidebarMenu.vue'
import BmImageList from '../components/molecules/BmImageList.vue'
import BmClassesPage from '../components/organisms/BmClassesPage.vue'
import BmPolygonList from '../components/molecules/BmPolygonList.vue'
import { useTokenStore } from '../store.js'
import BmCanva from '../components/atoms/BmCanva.vue'
import BmPopup from '../components/atoms/BmPopup.vue'
import BmButton from '../components/atoms/BmButton.vue'
import BmInput from '../components/atoms/BmInput.vue'
import BmClassSelection from '../components/atoms/BmClassSelection.vue'

const props = defineProps({
  projectID: String,
})

const selectedTab = ref('images')
const selectedTool = ref('select')
const selectedImage = ref('')
const selectedUrl = ref('')

const projectName = ref('')

const selectedOption = ref(null)

const polygonCreateName = ref('')

const polygonEditShow = ref(false)
const polygonEditId = ref('')
const polygonEditName = ref('')
const polygonEditCoords = ref([])

const images = ref([])
const polygons = ref([])

const tokenStore = useTokenStore()

const polygonCreateShow = ref(false)

  fetch(`http://localhost:8000/api/project/${props.projectID}`, {
    method: "GET", // Or POST, PUT, DELETE, etc.
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON
  .then(data => {
    projectName.value = data.data[0].title
    console.log(data.data)
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });

fetch(`http://localhost:8000/api/project/${props.projectID}/images`, {
    method: "GET", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON
  .then(data => {
    images.value = data.data; 
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });



const handleTabChange = (tab) => {
  selectedTab.value = tab
  console.log(polygons)
}

const handleToolChange = (tool) => {
  selectedTool.value = tool
}

const handleImageSelect = (id) => {
  const image = images.value.find(image => image.id === id)
  selectedImage.value = id
  selectedUrl.value = image.url
  console.log('Selected image ID:', image.url) 

  fetch(`http://localhost:8000/api/project/image/${image.id}/polygons`, {
      method: "GET", // Or POST, PUT, DELETE, etc.
      headers: {
          "Authorization": `Bearer ${tokenStore.token}`,
          "Content-Type": "application/json" // Example for JSON data
      },
    })
    .then(response => response.json()) // Parse the response as JSON
    .then(data => {
      polygons.value = data.data; 
    })
    .catch(error => {
      console.error('Error fetching data:', error);
    });
}

const handleZoomIn = () => {
  // TODO: Implement zoom in
}

const handleZoomOut = () => {
  // TODO: Implement zoom out
}


const handleEditPolygon = (polygon) => {
  polygonEditShow.value = true
  polygonEditId.value = polygon.id
  polygonEditName.value = polygon.name
  selectedOption.value = polygon.classID
  polygonEditCoords.value = polygon.coordinates
}

const handleDeletePolygon = (id) => {
  fetch(`http://localhost:8000/api/project/image/${selectedImage.value}/polygon/${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    }
  })
  .then(response => {
    if (response.ok) {
      polygons.value = polygons.value.filter(p => p.id !== id)
    }
  })
  .catch(error => {
    console.error('Error deleting polygon:', error);
  });
}

function createPolygon() {
  fetch(`http://localhost:8000/api/project/image/${selectedImage.value}/polygon/`, {
      method: "POST", // Or POST, PUT, DELETE, etc.
      headers: {
          "Authorization": `Bearer ${tokenStore.token}`,
          "Content-Type": "application/json" // Example for JSON data
      },
      body: JSON.stringify({
        name: polygonCreateName.value,
        classID: selectedOption.value.id,
        coordinates: [[0, 0], [0, 10], [10, 10], [0, 0]]
      })
    })
    .then(response => response.json()) // Parse the response as JSON      
    .then(data => {
      polygons.value.push(data.data[0]); 
      polygonCreateShow.value = false
    })
    .catch(error => {
      console.error('Error fetching data:', error);
    }); 
}

function openPolygon() {
  polygonCreateShow.value = true
  console.log('Open polygon for image ID:', selectedImage.value)
}


function updatePolygon() {
  console.log(polygonEditCoords.value)
  fetch(`http://localhost:8000/api/project/image/${selectedImage.value}/polygon/${polygonEditId.value}`, {
    method: "PUT",
    body: JSON.stringify({
        name: polygonEditName.value,
        classID: selectedOption.value.id,
        coordinates: polygonEditCoords.value
    }),
    headers: {  
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    },
  })
  .then(response => response.json())
  .then(data => {
    console.log(data)
    polygonEditShow.value = false; // Close the popup
    polygonEditId.value = null; // Clear the class id
    polygonEditName.value = ''; // Clear the class name input
  })
  .catch(error => {
    console.error('Error editing class:', error)
  })
}
</script>

<template>
<div class="font-display bg-background-light dark:bg-background-dark text-gray-800 dark:text-gray-200">
<div class="flex flex-col h-screen w-full">
<header class="flex items-center justify-between whitespace-nowrap border-b border-solid border-white/10 px-6 py-3 bg-background-dark flex-shrink-0">
<div class="flex items-center gap-4 text-white">
<div class="size-5 text-primary">
<svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
<path d="M24 4C25.7818 14.2173 33.7827 22.2182 44 24C33.7827 25.7818 25.7818 33.7827 24 44C22.2182 33.7827 14.2173 25.7818 4 24C14.2173 22.2182 22.2182 14.2173 24 4Z" fill="currentColor"></path>
</svg>
</div>
<h2 class="text-white text-lg font-bold leading-tight tracking-tight">Bug Marker</h2>
<div class="flex items-center gap-2 text-sm text-gray-400">
<span class="material-symbols-outlined text-base">chevron_right</span>
<span>{{ projectName }}</span>
</div>
</div>
<div class="flex gap-3">
<button class="flex max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 bg-white/10 text-white gap-2 text-sm font-bold leading-normal tracking-[0.015em] min-w-0 px-2.5 hover:bg-white/20 transition-colors">
<span class="material-symbols-outlined !text-2xl">account_circle</span>
</button>
</div>
</header>
<div class="flex flex-1 overflow-hidden">
<aside class="w-1/4 max-w-[360px] min-w-[300px] flex-shrink-0 bg-background-dark flex flex-col">
<div class="p-4 border-b border-white/10">
<BmSidebarMenu :selected-tab="selectedTab" @tab-change="handleTabChange" />
</div>
<BmImageList 
  v-if="selectedTab === 'images'"
  :projectID="props.projectID" 
  :images="images" 
  :selected-image="selectedImage" 
  @select-image="handleImageSelect" 
/>
<BmPolygonList
  v-if="selectedTab === 'polygons'"
  :current-image="selectedImage"
  :polygons="polygons"
  @new="openPolygon"
  @edit="handleEditPolygon"
  @delete="handleDeletePolygon"
  @add-point="handleAddPoint"
  :projectID="projectID" 
/>
</aside>
<main class="flex-1 w-full flex flex-col bg-[#111111] overflow-hidden border-l border-white/10">
  <BmClassesPage :projectId="props.projectID" v-if="selectedTab === 'classes'" />
  <BmCanva :projectID="projectID" :imageURL="selectedUrl" :polygons="polygons" v-if="selectedTab !== 'classes'"/>
</main>

</div>
</div>
</div>
<BmPopup :show="polygonCreateShow" title="Criar Polígono">
  <button
    @click="polygonCreateShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
    >
    ✕
  </button>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome do Polígono</p>
      <BmInput
        placeholder="digite o nome do polígono"
        v-model="polygonCreateName"
      />
  </label>
  </div>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">classe de marcação</p>

    <BmClassSelection
      v-model="selectedOption"
      :projectID="props.projectID"
      />

  </label>
  </div>
  <BmButton @click="createPolygon" type="submit">Editar polígono</BmButton>
</BmPopup>
<BmPopup :show="polygonEditShow" title="Criar Polígono">
  <button
    @click="polygonEditShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
    >
    ✕
  </button>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome do Polígono</p>
      <BmInput
        placeholder="digite o nome do polígono"
        v-model="polygonEditName"
      />
  </label>
  </div>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Classe de marcação</p>

    <BmClassSelection
      v-model="selectedOption"
      :projectID="props.projectID"
      />

  </label>
  </div>
  <BmButton @click="updatePolygon" type="submit">Modificar polígono</BmButton>
</BmPopup>
</template>