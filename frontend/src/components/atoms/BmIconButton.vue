<template>
  <button 
    class="flex items-center justify-center size-8 rounded-full bg-white/10 hover:bg-white/20 text-white transition-colors"
    :title="title"
    @click="$emit('click')"
  >
    <span class="material-symbols-outlined text-base">{{ icon }}</span>
  </button>
</template>

<script setup>
defineProps({
  icon: {
    type: String,
    required: true
  },
  title: {
    type: String,
    default: ''
  }
})

defineEmits(['click'])
</script>