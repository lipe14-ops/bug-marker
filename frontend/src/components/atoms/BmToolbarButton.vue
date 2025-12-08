<template>
  <button 
    :class="[
      'p-2.5 rounded-md transition-colors',
      active 
        ? 'text-primary bg-primary/20' 
        : 'text-gray-300 hover:bg-white/10 hover:text-white'
    ]"
    :title="title"
    @click="$emit('click')"
  >
    <span class="material-symbols-outlined !text-xl">{{ icon }}</span>
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
  },
  active: {
    type: Boolean,
    default: false
  }
})

defineEmits(['click'])
</script>