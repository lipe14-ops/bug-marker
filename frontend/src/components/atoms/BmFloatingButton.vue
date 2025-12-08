<template>
  <button 
    class="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-xl h-12 px-6 bg-primary text-background-dark text-base font-bold leading-normal tracking-[0.015em] hover:opacity-90 transition-opacity gap-2"
    @click="$emit('click')"
  >
    <span class="material-symbols-outlined">add</span>
    <span class="truncate">{{ text }}</span>
  </button>
</template>

<script setup>
defineProps({
  text: {
    type: String,
    default: ''
  }
})

defineEmits(['click'])
</script>