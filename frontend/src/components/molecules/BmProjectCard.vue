<template>
  <div class="h-full group relative flex flex-col justify-between gap-4 p-4 rounded-lg bg-black/50 border border-white/10 overflow-hidden transition-all duration-300 hover:border-primary/80 hover:shadow-lg hover:shadow-primary/20 hover:-translate-y-1">
    <div class="flex flex-col gap-1">
      <p class="text-white text-base font-medium leading-normal truncate">{{ title }}</p>
      <p class="text-white/60 text-sm font-normal leading-normal line-clamp-2">{{ description }}</p>
      <p class="text-white/60 text-xs font-normal leading-normal">Criado em: {{ createdAt }}</p>
      <p class="text-white/60 text-xs font-normal leading-normal">{{ imagesCount }} imagens | {{ classesCount }} {{ classesCount === 1 ? 'classe' : 'classes' }}</p>
    </div>
    <div class="flex items-center justify-end gap-1">
      <BmIconButton 
        icon="arrow_outward"
        title="Acessar Projeto"
        @click="$emit('enter')"
      />
      <BmIconButton 
        icon="edit"
        title="Editar Projeto"
        @click="$emit('edit')"
      />
      <BmIconButton
        icon="delete"
        title="Excluir Projeto"
        @click="$emit('delete')"
      />
    </div>
  </div>
</template>

<script setup>
import BmIconButton from '../atoms/BmIconButton.vue'

defineProps({
  title: {
    type: String,
    required: true
  },
  description: {
    type: String,
    default: ''
  },
  createdAt: {
    type: String,
    required: true
  },
  imagesCount: {
    type: Number,
    required: true
  },
  classesCount: {
    type: Number,
    required: true
  }
})

defineEmits(['enter', 'edit', 'delete'])
</script>