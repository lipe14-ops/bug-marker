<template>
  <div class="bg-white/5 p-4 rounded-lg">
    <h3 class="text-lg font-bold text-white mb-4">Criar Nova Classe</h3>
    <div class="flex flex-col sm:flex-row items-start gap-4">
      <div class="flex-1 w-full">
        <BmInput
          id="class-name"
          label="Nome da Classe"
          v-model="className"
          placeholder="Ex: Antracnose"
        />
      </div>
      <div class="flex-shrink-0">
        <BmInput
          id="class-color"
          label="Cor"
          v-model="classColor"
          type="color"
        />
      </div>
      <div class="self-end pt-2 sm:pt-0">
        <button 
          class="flex items-center justify-center gap-2 w-full sm:w-auto min-w-[84px] cursor-pointer overflow-hidden rounded-lg h-10 px-5 bg-white/10 text-white text-sm font-bold leading-normal hover:bg-white/20 transition-colors"
          @click="handleCreate"
        >
          <span>Criar</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import BmInput from '../atoms/BmInput.vue'

const className = ref('')
const classColor = ref('#06F906')

const emit = defineEmits(['create'])

const handleCreate = () => {
  if (!className.value) return

  emit('create', {
    name: className.value,
    color: classColor.value
  })

  className.value = ''
  classColor.value = '#06F906'
}
</script>