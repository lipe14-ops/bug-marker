<template>
  <div class=" p-8" v-if="show">
    <!-- Fundo escurecido -->
    <div
      class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
      @click.self="show = false"
    >
      <!-- Caixa do popup -->
      <div class="bg-background-light dark:bg-background-dark rounded-lg shadow-lg p-6 w-96 relative">
        <h2 class="text-xl text-black dark:text-white font-semibold mb-4">{{ title }}</h2>
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup>

const props = defineProps({
    show: Boolean,
    title: String,
})
</script>
