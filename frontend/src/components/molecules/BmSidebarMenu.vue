<template>
  <div class="flex flex-col gap-2">
    <BmSidebarItem
      icon="image"
      label="Imagens"
      :active="selectedTab === 'images'"
      @click="$emit('tab-change', 'images')"
    />
    <BmSidebarItem
      icon="label"
      label="Classes de Marcação"
      :active="selectedTab === 'classes'"
      @click="$emit('tab-change', 'classes')"
    />
    <BmSidebarItem
      icon="hexagon"
      label="Polígonos da Imagem"
      :active="selectedTab === 'polygons'"
      @click="$emit('tab-change', 'polygons')"
    />
  </div>
</template>

<script setup>
import BmSidebarItem from '../atoms/BmSidebarItem.vue'

defineProps({
  selectedTab: {
    type: String,
    default: 'images'
  }
})

defineEmits(['tab-change'])
</script>