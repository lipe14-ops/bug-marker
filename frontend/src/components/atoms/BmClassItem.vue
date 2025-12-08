<template>
  <li class="flex items-center justify-between bg-white/5 p-3 rounded-lg hover:bg-white/10 transition-colors">
    <div class="flex items-center gap-3">
      <div 
        class="w-4 h-4 rounded-full" 
        :style="{ backgroundColor: color }"
      ></div>
      <p class="text-white font-medium">{{ name }}</p>
    </div>
    <div class="flex items-center gap-2">
      <button 
        class="p-2 rounded-md text-gray-300 hover:bg-white/20 hover:text-white transition-colors" 
        title="Editar"
        @click="$emit('edit')"
      >
        <span class="material-symbols-outlined !text-xl">edit</span>
      </button>
      <button 
        class="p-2 rounded-md text-gray-300 hover:bg-red-500/20 hover:text-red-400 transition-colors" 
        title="Remover"
        @click="$emit('delete')"
      >
        <span class="material-symbols-outlined !text-xl">delete</span>
      </button>
    </div>
  </li>
</template>

<script setup>
defineProps({
  name: {
    type: String,
    required: true
  },
  color: {
    type: String,
    required: true
  }
})

defineEmits(['edit', 'delete'])
</script>