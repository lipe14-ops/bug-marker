<template>
  <details 
    class="group bg-white/5 rounded-lg border border-white/10 hover:border-white/20 transition-colors duration-300 mt-4"
    :class="[{ 'shadow-lg': isEditing }, { 'open:border-primary/50': !isOpen }]"
    :style="isEditing ? { borderColor: '#08E308', boxShadow: '0 0 0 2px #08E308' } : {}"
    :open="isOpen"
  >
    <summary class="p-4 list-none flex justify-between items-center cursor-pointer">
      <div class="flex flex-col">
        <h4 class="font-bold text-primary">{{ name }}</h4>
        <div class="mt-1 flex items-center gap-2">
          <BmColorIndicator :color="classColor" />
          <p class="text-xs text-gray-300">{{ className }}</p>
        </div>
      </div>
      <div class="flex items-center gap-1">
        <BmPolygonActionButton
          icon="edit"
          title="Editar polígono"
          @click="$emit('edit')"
        />
        <BmPolygonActionButton
          icon="add"
          title="Adicionar ponto"
          @click="$emit('add-point', id)"
        />
        <BmPolygonActionButton
          icon="delete"
          title="Remover polígono"
          variant="danger"
          @click="$emit('delete')"
        />
        <span class="material-symbols-outlined !text-2xl text-gray-400 transition-transform duration-300 group-open:rotate-180">expand_more</span>
      </div>
    </summary>
    <div class="px-4 pb-4">
      <div class="max-h-[200px] overflow-y-auto points-table-container">
        <table class="w-full text-xs text-left">
          <thead class="text-gray-400 sticky top-0 bg-white/5 z-10">
            <tr>
              <th class="font-semibold py-2 w-1/3">Ponto</th>
              <th class="font-semibold py-2 w-1/3 text-center">X</th>
              <th class="font-semibold py-2 w-1/3 text-center">Y</th>
            </tr>
          </thead>
          <tbody class="text-gray-300 font-mono">
            <BmPolygonPoint
              v-for="(point, index) in points"
              :key="index"
              :point="`#${index + 1}`"
              :x="point[0]"
              :y="point[1]"
            />
          </tbody>
        </table>
      </div>
    </div>
  </details>
</template>

<script setup>
import BmPolygonPoint from './BmPolygonPoint.vue'
import BmColorIndicator from './BmColorIndicator.vue'
import BmPolygonActionButton from './BmPolygonActionButton.vue'

defineProps({
  id: {
    type: [String, Number],
    required: false
  },
  name: {
    type: String,
    required: true
  },
  className: {
    type: String,
    required: true
  },
  classColor: {
    type: String,
    required: true
  },
  points: {
    type: Array,
    required: true
  },
  isOpen: {
    type: Boolean,
    default: false
  },
  isEditing: {
    type: Boolean,
    default: false
  }
})

defineEmits(['edit', 'delete', 'add-point'])
</script>