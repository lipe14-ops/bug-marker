<template>
  <button 
    :class="[
      'flex w-full items-center justify-center rounded-lg px-6 py-4',
      'text-base font-bold text-background-dark transition-all',
      'bg-primary hover:bg-opacity-90 h-14'
    ]"
    :type="type"
  >
    <slot></slot>
  </button>
</template>

<script setup>
defineProps({
  type: {
    type: String,
    default: 'button'
  }
})
</script>