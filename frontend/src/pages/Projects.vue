<style>
  .glassmorphism {
      background: rgba(4, 30, 4, 0.4);
      -webkit-backdrop-filter: blur(10px);
      backdrop-filter: blur(10px);
      border: 1px solid rgba(255, 255, 255, 0.1);
  }
</style>

<script setup>
import BmProjectCard from '../components/molecules/BmProjectCard.vue'
import BmIconButton from '../components/atoms/BmIconButton.vue'
import BmInput from '../components/atoms/BmInput.vue'
import BmFloatingButton from '../components/atoms/BmFloatingButton.vue'
import BmPopup from '../components/atoms/BmPopup.vue'
import { ref } from 'vue'
import { useTokenStore } from '../store.js'
import BmButton from '../components/atoms/BmButton.vue'
import { router } from '../router.js'

const searchQuery = ref('')

const projects = ref([])

const props = defineProps({
  ownerID: String
})

const projectCreateShow = ref(false)
const projectCreateName = ref('')
const projectCreateDescription = ref('')

const projectUpdateShow = ref(false)
const projectUpdateId = ref(null)
const projectUpdateName = ref('')
const projectUpdateDescription = ref('')

const tokenStore = useTokenStore()

fetch(`http://localhost:8000/api/projects/owner/${props.ownerID}`, {
    method: "GET", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON
  .then(data => {
    console.log(data); // Work with the fetched data
    projects.value = data.data; // Update the projects ref with the fetched data
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });

const handleEdit = (id) => {
  console.log('Edit project:', id)

  const project = projects.value.find(project => project.id === id)
  projectUpdateName.value = project.title
  projectUpdateDescription.value = project.description
  projectUpdateShow.value = true
  projectUpdateId.value = id
}

function atualizarProjeto() {
  fetch(`http://localhost:8000/api/project/${projectUpdateId.value}`, {
    method: "PUT", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
    body: JSON.stringify({
      id: projectUpdateId.value,
      title: projectUpdateName.value,
      description: projectUpdateDescription.value,
      owner: props.ownerID
    })
  })
  .then(response => response.json()) // Parse the response as JSON
  .then(data => {
    projects.value = projects.value.map(project => {
      if (project.id === projectUpdateId.value) {
        return data.data[0]; // Update the project with the new data
      }
      return project;
    });
    projectUpdateShow.value = false; // Close the popup
    projectUpdateId.value = null; // Clear the project id
    projectUpdateName.value = ''; // Clear the project name input
    projectUpdateDescription.value = ''; // Clear the project description input
  }).catch(error => {
    console.error('Error fetching data:', error); 
  });
}

const handleDelete = (id) => {
  let result = confirm('Tem certeza que deseja excluir este projeto? Esta ação não pode ser desfeita.')
  
  if (!result) {
    return
  }

  fetch(`http://localhost:8000/api/project/${id}`, {
    method: "DELETE", // Or POST, PUT, DELETE, etc.
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON
  .then(data => {
    projects.value = projects.value.filter(project => project.id !== id);
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });
}

function adicionarProjeto() {
    fetch('http://localhost:8000/api/project', {
      method: "POST", // Or POST, PUT, DELETE, etc.
      body: JSON.stringify({
        title: projectCreateName.value,
        description: projectCreateDescription.value,
        owner: props.ownerID
      }),
      headers: {
          "Authorization": `Bearer ${tokenStore.token}`,
          "Content-Type": "application/json" // Example for JSON data
      },
    })
    .then(response => response.json()) // Parse the response as JSON
    .then(data => {
      console.log(data); // Work with the fetched data
      projects.value.push(data.data[0]); // Add the new project to the projects list
      projectCreateShow.value = false; // Close the popup
      projectCreateName.value = ''; // Clear the project name input
      projectCreateDescription.value = ''; // Clear the project description input
    })
    .catch(error => {
      console.error('Error fetching data:', error);
    });
}

function pushProject(id) {
  console.log('Entering project:', id)
  router.push({ path: `/project/${id}` });
}

</script>

<template>
<div class="bg-black font-display text-[#E0E0E0]">
<div class="relative min-h-screen w-full flex flex-col items-center">
<div class="absolute top-0 left-0 w-96 h-96 bg-primary/20 rounded-full filter blur-3xl opacity-30 animate-pulse"></div>
<div class="absolute bottom-0 right-0 w-96 h-96 bg-primary/20 rounded-full filter blur-3xl opacity-30 animate-pulse delay-1000"></div>
<div class="w-full max-w-7xl px-4 sm:px-6 lg:px-8 z-10">
<header class="flex items-center justify-between whitespace-nowrap border-b border-solid border-white/10 py-4">
<div class="flex items-center gap-4 text-white">
<div class="size-6 text-primary">
<svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
<path d="M24 4C25.7818 14.2173 33.7827 22.2182 44 24C33.7827 25.7818 25.7818 33.7827 24 44C22.2182 33.7827 14.2173 25.7818 4 24C14.2173 22.2182 22.2182 14.2173 24 4Z" fill="currentColor"></path>
</svg>
</div>
<h2 class="text-white text-xl font-bold tracking-[-0.015em]">Bug Marker</h2>
</div>
<div class="flex items-center gap-4">
<BmIconButton
  icon="person"
  title="Perfil do Usuário"
  class="h-10 w-10"
/>
</div>
</header>
<main class="py-10">
<div class="flex flex-wrap justify-between items-center gap-4 mb-8">
<div class="flex flex-col gap-2">
<p class="text-white text-4xl font-black leading-tight tracking-[-0.033em]">Meus Projetos</p>
<p class="text-white/60 text-base font-normal leading-normal">Gerencie, edite e crie novos projetos de análise.</p>
</div>
<BmFloatingButton text="Adicionar Novo Projeto" @click="projectCreateShow = true" />
</div>
<div class="mb-8">
  <label class="flex flex-col w-full">
    <div class="rounded-xl border border-white/10 bg-black/50">
      <BmInput
        v-model="searchQuery"
        placeholder="Buscar por nome do projeto..."
        icon="search"
        darkMode
        containerClass="h-12"
        className="h-full px-4"
      />
    </div>
  </label>
</div>
<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
    <BmProjectCard
    v-for="project in projects"
    :key="project.id"
    :title="project.title"
    :created-at="project.createdAt"
    :images-count="project.imagesCount"
    :classes-count="project.classesCount"
    :description="project.description"
    @enter="pushProject(project.id)"
    @edit="handleEdit(project.id)"
    @delete="handleDelete(project.id)"
  />
</div>
</main>
</div>
</div>
</div>
<BmPopup :show="projectCreateShow" title="Novo projeto">
  <button
    @click="projectCreateShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
    >
    ✕
  </button>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome do projeto</p>
    <div class="rounded-xl border border-white/10 bg-black/50">
      <BmInput
        placeholder="digite o nome"
        v-model="projectCreateName"
      />
    </div>
  </label>
  </div>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Descrição do projeto</p>
    <div class="rounded-xl border border-white/10 bg-black/50">
      <BmInput
        placeholder="digite a descrição"
        v-model="projectCreateDescription"
      />
    </div>
  </label>
  </div>
  <BmButton @click="adicionarProjeto" type="submit">Adicionar projeto</BmButton>
</BmPopup>

<BmPopup :show="projectUpdateShow" title="Atualizar projeto">
  <button
    @click="projectUpdateShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
    >
    ✕
  </button>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome do projeto</p>
    <div class="rounded-xl border border-white/10 bg-black/50">
      <BmInput
        placeholder="digite o nome"
        v-model="projectUpdateName"
      />
    </div>
  </label>
  </div>
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Descrição do projeto</p>
    <div class="rounded-xl border border-white/10 bg-black/50">
      <BmInput
        placeholder="digite a descrição"
        v-model="projectUpdateDescription"
      />
    </div>
  </label>
  </div>
  <BmButton @click="atualizarProjeto(projectUpdateId)" type="submit">Atualizar projeto</BmButton>
</BmPopup>
</template>