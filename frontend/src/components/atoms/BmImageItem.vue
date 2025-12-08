<template>
  <div class="flex flex-row justify-between">
    <span 
      :class="[
        'flex items-center justify-between p-3 rounded-lg w-60',
        selected 
          ? 'border-2 border-primary ring-2 ring-primary/30 bg-primary/10' 
          : 'cursor-pointer hover:bg-white/10 transition-colors'
      ]"
      @click="$emit('click')"
    >
      <!-- Nome do arquivo à esquerda -->
      <p class="text-white text-xs font-semibold truncate">{{ name }}</p>
  
      <!-- Botão de deletar à direita -->
    </span>
    <div class="flex flex-row gap-1">
      <BmIconButton
        icon="edit"
        title="Editar Imagem"
        @click="$emit('edit', imageId)"
      />
      <BmIconButton
        icon="delete"
        title="Excluir Imagem"
        @click="$emit('delete', imageId)"
      />
    </div>
  </div>
</template>


<script setup>
import BmIconButton from './BmIconButton.vue';

defineProps({
  name: {
    type: String,
    required: true
  },
  imageId: {
    type: String,
    required: true
  },
  selected: {
    type: Boolean,
    default: false
  }
})

defineEmits(['edit', 'delete'])
</script>