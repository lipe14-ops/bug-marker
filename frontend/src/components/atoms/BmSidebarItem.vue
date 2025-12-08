<template>
  <div 
    :class="[
      'flex items-center gap-3 px-3 py-2.5 rounded-lg cursor-pointer',
      active ? 'bg-primary/20' : 'hover:bg-white/10 transition-colors'
    ]"
    @click="$emit('click')"
  >
    <span 
      class="material-symbols-outlined"
      :class="active ? 'text-primary' : 'text-gray-400'"
    >{{ icon }}</span>
    <p 
      :class="[
        'text-sm leading-normal',
        active ? 'text-primary font-bold' : 'text-gray-300 font-medium'
      ]"
    >{{ label }}</p>
  </div>
</template>

<script setup>
defineProps({
  icon: {
    type: String,
    required: true
  },
  label: {
    type: String,
    required: true
  },
  active: {
    type: Boolean,
    default: false
  }
})

defineEmits(['click'])
</script>