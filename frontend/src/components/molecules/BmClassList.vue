<template>
  <div class="space-y-3">
    <h3 class="text-lg font-bold text-white">Classes Existentes ({{ classes.length }})</h3>
    <ul class="space-y-3">
      <BmClassItem
        v-for="item in classes"
        :key="item.id"
        :id="item.id"
        :name="item.name"
        :color="item.color"
        @edit="$emit('edit', item)"
        @delete="$emit('delete', item.id)"
      />
    </ul>
  </div>
</template>

<script setup>
import BmClassItem from '../atoms/BmClassItem.vue'

defineProps({
  classes: {
    type: Array,
    required: true
  }
})

defineEmits(['edit', 'delete'])
</script>