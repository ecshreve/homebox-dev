<template>
  <div class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <input type="color" v-model="value" class="input input-primary w-full" />
  </div>
</template>

  <script setup lang="ts">
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
  </script>

  