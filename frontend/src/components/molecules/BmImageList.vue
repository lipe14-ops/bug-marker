<style>
.image-list-scroll {
  scrollbar-width: thin;
  scrollbar-color: rgba(255, 255, 255, 0.2) rgba(255, 255, 255, 0.05);
}

.image-list-scroll::-webkit-scrollbar {
  width: 6px;
}

.image-list-scroll::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
}

.image-list-scroll::-webkit-scrollbar-thumb {
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 3px;
}

.image-list-scroll::-webkit-scrollbar-thumb:hover {
  background-color: rgba(255, 255, 255, 0.3);
}
</style>

<template>
  <div class="flex flex-col min-h-0 ">
    <div class="p-4 flex-shrink-0 border-b border-white/10">
      <div class="flex justify-between items-center mb-4">
        <div class="space-y-1">
          <h3 class="text-white font-bold text-lg">Imagens</h3>
          <p class="text-gray-400 text-xs">{{ images.length }} imagens encontradas</p>
        </div>
        <button 
          class="flex items-center justify-center gap-2 rounded-lg h-9 px-4 bg-primary text-background-dark text-sm font-bold hover:bg-primary/90 transition-colors"
          @click="imageCreateShow = true"
        >
          <span class="material-symbols-outlined !text-xl">add</span>
          <span>Nova Imagem</span>
        </button>
      </div>
      </div>
    <div class="flex-1 overflow-y-auto image-list-scroll">
      <div class="p-4 space-y-2">
        <BmImageItem
          v-for="image in images"
          :imageId="image.id"
          :name="image.filename"
          :selected="image.id === selectedImage"
          @edit="openEditPopup"
          @delete="deleteImage"
          @click="$emit('select-image', image.id)"
        />
      </div>
    </div>
  </div>
<BmPopup :show="imageCreateShow" title="Nova imagem">
  <button
    @click="imageCreateShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
  >
    ✕
  </button>
  
  <div class="mb-6">
 
    
    <!-- Input de arquivo estilizado -->
    <div class="relative border-2 border-dashed border-border-dark hover:border-primary rounded-xl transition-all duration-300 bg-background-dark/30 overflow-hidden group">
      <input
        type="file"
        @change="handleFileChange"
        accept="image/*"
        class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10"
      />
      
      <!-- Área de upload -->
      <div class="flex flex-col items-center justify-center py-12 px-6 group-hover:bg-primary/5 transition-colors">
        <div class="w-16 h-16 bg-primary/10 rounded-full flex items-center justify-center mb-4 group-hover:bg-primary/20 transition-colors">
          <svg class="w-8 h-8 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
          </svg>
        </div>
        <p class="text-text-dark text-lg font-semibold mb-2">
          Clique para fazer upload
        </p>
        <p class="text-subtle-text-dark text-sm mb-1">
          ou arraste e solte uma imagem aqui
        </p>
        <p class="text-subtle-text-dark/60 text-xs">
          PNG, JPG, GIF até 10MB
        </p>
      </div>
    </div>
  </div>
  
  <BmButton @click="uploadImage" type="submit" class="w-full">
    Adicionar imagem
  </BmButton>
</BmPopup>

<BmPopup :show="imageUploadShow" title="Editar imagem">
  <button
    @click="imageUploadShow = false"
    class="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
  >
    ✕
  </button>
  
  <div class="mb-4">
  <label class="flex flex-col w-full">
    <p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Nome do projeto</p>
      <BmInput
        v-model="imageUploadName"
        placeholder="Digite o novo nome da imagem"
      />
    </label>
  </div>

  <BmButton @click="updateImage" type="submit" class="w-full">
    Salvar alterações
  </BmButton>
</BmPopup>
</template>

<script setup>
import BmImageItem from '../atoms/BmImageItem.vue'
import BmPopup from '../atoms/BmPopup.vue'
import { useTokenStore } from '../../store.js'
import { ref  } from 'vue'
import BmButton from '../atoms/BmButton.vue'
import BmInput from '../atoms/BmInput.vue'

const imageCreateShow = ref(false)
const tokenStore = useTokenStore()

const file = ref(null)
const preview = ref(null)

const props = defineProps({
  selectedImage: {
    type: String,
    default: ''
  },
  projectID: {
    type: String,
    default: ''
  }
})

const images = ref([])

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


// Captura o arquivo selecionado
function handleFileChange(event) {
  const selectedFile = event.target.files[0]
  if (!selectedFile) return

  file.value = selectedFile
  preview.value = URL.createObjectURL(selectedFile) // preview da imagem
}

// Faz upload usando fetch
function uploadImage() {
  if (!file.value) return

  const formData = new FormData()
  formData.append('document', file.value)

  fetch(`http://localhost:8000/api/project/${props.projectID}/image`, {
    method: 'POST',
    body: formData,
      headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
    },
  }).then(response => response.json())
    .then(data => {
      console.log(data)
      imageCreateShow.value = false // Fecha o popup
      file.value = null // Limpa o arquivo selecionado
      preview.value = null // Limpa o preview
    })
    .catch(error => {
      console.error('Erro ao enviar imagem:', error)
    })

}

function deleteImage(imageId) {
  if (!confirm('Tem certeza que deseja deletar esta imagem?')) return

  fetch(`http://localhost:8000/api/project/image/${imageId}`, {
    method: 'DELETE',
    headers: {
      "Authorization": `Bearer ${tokenStore.token}`,
      "Content-Type": "application/json"
    },
  }).then(response => response.json())
    .then(data => {
      images.value = images.value.filter(image => image.id !== imageId)
    })    
    .catch(error => {
      console.error('Erro ao deletar imagem:', error)
    })
} 

const imageUploadShow = ref(false)
const imageUploadName = ref('')
const imageUploadId = ref('')

function openEditPopup(imageId) {
  imageUploadId.value = imageId
  imageUploadShow.value = true
}

function updateImage() {
  // Lógica para atualizar a imagem (a ser implementada)
  console.log('Atualizar imagem com ID:', imageUploadId.value)
  fetch(`http://localhost:8000/api/project/image/${imageUploadId.value}`, {
    method: "PUT", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
    body: JSON.stringify({
      filename: imageUploadName.value
    })
  }).then(response => response.json()) // Parse the response as JSON
  .then(data => {
    imageUploadShow.value = false // Fecha o popup
    imageUploadName.value = '' // Limpa o nome da imagem
    images.value = images.value.map(image => {
      if (image.id === imageUploadId.value) {
        return data.data[0]; // Update the image with the new data
      }
      return image;
    });
    imageUploadId.value = '' // Limpa o id da imagem
  })  
  .catch(error => {
    console.error('Error updating image:', error);
  }); 

} 

defineEmits(['select-image'])
</script>