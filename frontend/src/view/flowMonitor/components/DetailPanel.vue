<template>
  <div>
    <!-- Panel content with transition -->
    <transition :name="transitionName" :css="!!transitionName">
      <div v-show="localVisible" ref="panel">
        <!-- Panel header -->
        <div v-if="title || $slots.header">
          <slot name="header">
          </slot>
        </div>

        <!-- Panel content -->
        <div>
          <slot></slot>
        </div>

        <!-- Panel footer with toggle button -->
        <div v-if="showToggle">
          <slot name="footer"></slot>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'

const props = defineProps({
  // Visibility control
  visible: {
    type: Boolean,
    default: false
  },

  // Trigger button config
  showTrigger: {
    type: Boolean,
    default: true
  },
  triggerType: {
    type: String,
    default: 'primary'
  },
  showText: {
    type: String,
    default: 'Show Details'
  },
  hideText: {
    type: String,
    default: 'Hide Details'
  },

  // Panel content config
  title: String,
  showToggle: {
    type: Boolean,
    default: true
  },
  toggleType: {
    type: String,
    default: 'primary'
  },
  expandText: {
    type: String,
    default: 'Expand'
  },
  collapseText: {
    type: String,
    default: 'Collapse'
  },

  // Style config
  panelClass: [String, Array, Object],
  panelStyle: [String, Array, Object],
  shadow: {
    type: Boolean,
    default: true
  },
  rounded: {
    type: Boolean,
    default: true
  },
  border: {
    type: Boolean,
    default: true
  },
  bgColor: {
    type: String,
    default: 'bg-white'
  },

  // Animation config
  transitionName: {
    type: String,
    default: 'slide-fade'
  },
  scrollOnOpen: {
    type: Boolean,
    default: true
  },

  // State
  expanded: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits([
  'update:visible',
  'update:expanded',
  'toggle',
  'show',
  'hide'
])

const localVisible = ref(props.visible)
const localExpanded = ref(props.expanded)
const panel = ref(null)

const toggle = () => {
  localVisible.value = !localVisible.value
  emit('update:visible', localVisible.value)
  emit('toggle', localVisible.value)

  if (localVisible.value && props.scrollOnOpen) {
    nextTick(() => {
      panel.value?.scrollIntoView({ behavior: 'smooth' })
    })
  }
}

const toggleExpand = () => {
  localExpanded.value = !localExpanded.value
  emit('update:expanded', localExpanded.value)
}

watch(() => props.visible, (newVal) => {
  localVisible.value = newVal
  if (newVal) emit('show')
  else emit('hide')
})

watch(() => props.expanded, (newVal) => {
  localExpanded.value = newVal
})
</script>

<style scoped>

</style>
