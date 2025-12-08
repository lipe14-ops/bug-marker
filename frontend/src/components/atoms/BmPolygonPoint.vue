<template>
  <tr>
    <td class="py-1">{{ point }}</td>
    <td class="py-1 text-center">{{ x }}</td>
    <td class="py-1 text-center">{{ y }}</td>
  </tr>
</template>

<script setup>
defineProps({
  point: {
    type: String,
    required: true
  },
  x: {
    type: Number,
    required: true
  },
  y: {
    type: Number,
    required: true
  }
})
</script>