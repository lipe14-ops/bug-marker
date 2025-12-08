<template>
  <div 
    class="w-3 h-3 rounded-full border" 
    :style="{
      backgroundColor: color,
      borderColor: `${color}dd`
    }"
  ></div>
</template>

<script setup>
defineProps({
  color: {
    type: String,
    required: true
  }
})
</script>