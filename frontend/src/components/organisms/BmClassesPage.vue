<style>
.classes-scroll {
  scrollbar-width: thin;
  scrollbar-color: rgba(255, 255, 255, 0.2) rgba(255, 255, 255, 0.05);
}

.classes-scroll::-webkit-scrollbar {
  width: 6px;
}

.classes-scroll::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
}

.classes-scroll::-webkit-scrollbar-thumb {
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
}

.classes-scroll::-webkit-scrollbar-thumb:hover {
  background-color: rgba(255, 255, 255, 0.3);
}
</style>

<template>
  <main class="flex-1 flex flex-col bg-[#111111] overflow-hidden border-l border-white/10">
    <div class="flex-shrink-0 p-6 border-b border-white/10">
      <div class="max-w-4xl mx-auto">
        <div class="flex items-center justify-between">
          <h1 class="text-2xl font-bold text-white">Classes de Marcação</h1>
          <button 
            class="flex items-center justify-center gap-2 min-w-[84px] max-w-[480px] cursor-pointer overflow-hidden rounded-lg h-10 px-5 bg-primary text-background-dark text-sm font-bold leading-normal tracking-wide hover:bg-primary/90 transition-colors"
            @click="showForm = true"
          >
            <span class="material-symbols-outlined !text-xl">add_circle</span>
            <span>Adicionar Nova Classe</span>
          </button>
        </div>
      </div>
    </div>
    <div class="flex-1 min-h-0 overflow-y-auto classes-scroll">
      <div class="p-6 md:p-8 lg:p-12 space-y-6">
        <div class="max-w-4xl mx-auto space-y-6">
          <BmClassForm 
            v-if="showForm"
            @create="handleCreateClass"
          />
          <BmClassList
            :classes="classes"
            @edit="classShowPopup"
            @delete="handleDeleteClass"
          />
        </div>
      </div>
    </div>
  </main>
  <BmPopup
    :show="classUpdateShow"
    title="Editar Classe"
  >
    <button
    @click="classUpdateShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
    >
    ✕
  </button>
  <div class="mb-8">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome da classe</p>
    <BmInput
      v-model="classUpdateName"
      placeholder="digite o nome da classe"
    />
  </label> 
  </div>
<div class="mb-8">
  <label class="flex flex-col w-full">
      <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Cor da classe</p>
      <div class="flex-shrink-0">
        <BmInput
          id="class-color"
          label="Cor"
          v-model="classUpdateColor"
          type="color"
        />
    </div>
  </label> 
  </div>
  <div class="flex justify-end">
    <BmButton @click="handleEditClass" class="ml-2">Salvar</BmButton>
  </div>
  </BmPopup>
</template>

<script setup>
import { ref } from 'vue'
import BmClassForm from '../molecules/BmClassForm.vue'
import BmClassList from '../molecules/BmClassList.vue'
import BmPopup from '../atoms/BmPopup.vue'
import BmButton from '../atoms/BmButton.vue'
import BmInput from '../atoms/BmInput.vue'
import { useTokenStore } from '../../store.js'

const showForm = ref(false)
const classes = ref([])

const tokenStore = useTokenStore()

const props = defineProps({
  projectId: String
})

fetch(`http://localhost:8000/api/project/${props.projectId}/classes/`, {
  method: "GET",
  headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
  },
})
.then(response => response.json())
.then(data => {
  classes.value = data.data
})
.catch(error => {
  console.error('Error fetching classes:', error)
})

const handleCreateClass = (newClass) => {
  fetch(`http://localhost:8000/api/project/${props.projectId}/class/`, {
    method: "POST",
    body: JSON.stringify(newClass),
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    },
  })
  .then(response => response.json())
  .then(data => {
    classes.value.push(data.data[0])
  })
  .catch(error => {
    console.error('Error creating class:', error)
  })
}

const classUpdateShow = ref(false)
const classUpdateName = ref('')
const classUpdateId = ref(null)
const classUpdateColor = ref('')

function classShowPopup(classItem) {
  classUpdateId.value = classItem.id
  classUpdateColor.value = classItem.color
  classUpdateShow.value = true
}

const handleEditClass = () => {
  // TODO: Implement edit functionality
  console.log('Edit class:', classUpdateId.value)

  fetch(`http://localhost:8000/api/project/${props.projectId}/class`, {
    method: "PUT",
    body: JSON.stringify({
      id: classUpdateId.value,
      name: classUpdateName.value,
      color: classUpdateColor.value,
    }),
    headers: {  
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    },
  })
  .then(response => response.json())
  .then(data => {
    console.log(classes.value)
    classes.value = classes.value.map(item => { 
      if (item.id === classUpdateId.value) {
        return data.data[0]; // Update the class with the new data
      }
      return item;
    });
    classUpdateShow.value = false; // Close the popup
    classUpdateId.value = null; // Clear the class id
    classUpdateName.value = ''; // Clear the class name input
    classUpdateColor.value = ''; // Clear the class color input
  })
  .catch(error => {
    console.error('Error editing class:', error)
  })
}

const handleDeleteClass = (id) => {
  fetch(`http://localhost:8000/api/project/${props.projectId}/class/${id}`, {
    method: "DELETE",
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    },
  })
  .then(response => response.json())
  .then(data => {
    classes.value = classes.value.filter(item => item.id !== id)
  })
  .catch(error => {
    console.error('Error deleting class:', error)
  })
}
</script>