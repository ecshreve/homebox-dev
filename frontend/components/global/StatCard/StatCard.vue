<template>
  <div class="stats rounded-md bg-neutral shadow">
    <div class="stat space-y-1 p-3 text-center text-neutral-content">
      <div class="stat-title text-neutral-content">{{ title }}</div>
      <div class="stat-value text-2xl">
        <Currency v-if="type === 'currency'" :amount="value" />
        <template v-if="type === 'number'">{{ value }}</template>
      </div>
      <div v-if="subtitle" class="stat-desc">{{ subtitle }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import type { StatsFormat } from "./types";

  type Props = {
    title: string;
    value: number;
    subtitle?: string;
    type?: StatsFormat;
  };

  withDefaults(defineProps<Props>(), {
    type: "number",
    subtitle: undefined,
  });
</script>

<style></style>
