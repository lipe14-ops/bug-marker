<template>
  <div class="flex gap-1">
    <BmIconButton 
      icon="arrow_selector_tool" 
      title="Selecionar"
      :active="selectedTool === 'select'"
      @click="$emit('tool-change', 'select')"
    />
    <BmIconButton 
      icon="edit" 
      title="Desenhar Polígono"
      :active="selectedTool === 'draw'"
      @click="$emit('tool-change', 'draw')"
    />
    <BmIconButton 
      icon="zoom_in" 
      title="Zoom In"
      @click="$emit('zoom-in')"
    />
    <BmIconButton 
      icon="zoom_out" 
      title="Zoom Out"
      @click="$emit('zoom-out')"
    />
  </div>
</template>

<script setup>
import BmIconButton from '../atoms/BmIconButton.vue'

defineProps({
  selectedTool: {
    type: String,
    default: 'select'
  }
})

defineEmits(['tool-change', 'zoom-in', 'zoom-out'])
</script>