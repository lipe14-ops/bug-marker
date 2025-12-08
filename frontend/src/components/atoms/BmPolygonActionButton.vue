<template>
  <button 
    :class="[
      'p-1.5 rounded-md transition-colors',
      variant === 'danger' 
        ? 'text-gray-400 hover:text-red-500 hover:bg-red-500/10' 
        : 'text-gray-400 hover:text-white hover:bg-white/10'
    ]"
    :title="title"
    @click.stop="$emit('click')"
  >
    <span class="material-symbols-outlined !text-lg">{{ icon }}</span>
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
  variant: {
    type: String,
    default: 'default'
  }
})

defineEmits(['click'])
</script>