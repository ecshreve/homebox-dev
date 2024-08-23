<template>
  <div class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <div class="current-color" :style="{ backgroundColor: value }" @click="toggleOptions"></div>
    <div v-if="showOptions" class="color-modal">
      <div
        v-for="color in colors"
        :key="color"
        :style="{ backgroundColor: color }"
        class="color-option"
        :class="{ selected: value === color }"
        @click="selectColor(color)"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps({
  modelValue: {
    type: String,
    default: "#ff0000",
  },
  label: {
    type: String,
    default: "Select Color",
  },
});

const value = useVModel(props, "modelValue");
const colors = [
  "#83a598", // Bright Blue
  "#458588", // Blue
  "#689d6a", // Aqua
  "#8ec07c", // Soft Green
  "#b8bb26", // Bright Green
  "#98971a", // Green
  "#d79921", // Yellow
  "#fabd2f", // Bright Yellow
  "#fe8019", // Bright Orange
  "#fb4934", // Bright Red
  "#d3869b", // Pink
  "#b16286", // Purple
  "#928374", // Dark Gray
  "#a89984", // Gray
  "#bdae93", // Light4
  "#fbf1c7"  // Light7
];

const showOptions = ref(false);

function toggleOptions() {
  showOptions.value = !showOptions.value;
}

function selectColor(color: string) {
  value.value = color;
  showOptions.value = false;
}
</script>

<style scoped>
.color-options {
  display: flex;
  gap: 10px;
}

.color-option {
  width: 30px;
  height: 30px;
  cursor: pointer;
  border: 2px solid transparent;
}

.color-option.selected {
  border-color: black;
}

.current-color {
  width: 100%;
  height: 50px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  cursor: pointer;
}

.color-modal {
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  padding: 5px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}
</style>