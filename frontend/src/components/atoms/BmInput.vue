<template>
  <div class="flex w-full flex-1 items-stretch" :class="containerClass">
    <div v-if="icon" class="text-white/60 flex items-center justify-center pl-4">
      <span class="material-symbols-outlined">{{ icon }}</span>
    </div>
    <input
      :class="[
        'form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden',
        noRightBorder ? 'rounded-l-lg border-r-0 pr-2' : '',
        darkMode ? 'text-white focus:outline-0 focus:ring-0 border-none bg-transparent placeholder:text-white/60' : 'text-text-light dark:text-text-dark focus:outline-0 focus:ring-2 focus:ring-accent/50 border border-border-light dark:border-border-dark bg-white dark:bg-background-dark',
        className,
        'text-base font-normal leading-normal'
      ]"
      :type="type"
      :placeholder="placeholder"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
    />
  </div>
</template>

<script setup>
defineProps({
  type: {
    type: String,
    default: 'text'
  },
  placeholder: {
    type: String,
    default: ''
  },
  modelValue: {
    type: [String, Number],
    default: ''
  },
  noRightBorder: {
    type: Boolean,
    default: false
  },
  icon: {
    type: String,
    default: ''
  },
  darkMode: {
    type: Boolean,
    default: false
  },
  containerClass: {
    type: String,
    default: ''
  },
  className: {
    type: String,
    default: 'h-14 px-4'
  }
})

defineEmits(['update:modelValue'])
</script>