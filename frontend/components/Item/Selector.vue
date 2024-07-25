<template>
  <FormAutocomplete2 v-if="items" v-model="value" :items="items" display="name" label="Parent Item">
    <template #display="{ item, selected, active }">
      <div>
        <div class="flex w-full">
          {{ cast(item.value).name }}
          <span
            v-if="selected"
            :class="['absolute inset-y-0 right-0 flex  items-center pr-4', active ? 'text-white' : 'text-primary']"
          >
            <MdiCheck class="h-5 w-5" aria-hidden="true" />
          </span>
        </div>
      </div>
    </template>
  </FormAutocomplete2>
</template>

<script lang="ts" setup>
  import type { ItemSummary } from "~~/lib/api/types/data-contracts";
  import MdiCheck from "~icons/mdi/check";

  type Props = {
    modelValue?: ItemSummary | null;
  };

  // Cast the type of the item to a FlatTreeItem so we can get type "safety" in the template
  // Note that this does not actually change the type of the item, it just tells the compiler
  // that the type is FlatTreeItem. We must keep this in sync with the type of the items
  function cast(value: any): ItemSummary {
    return value as ItemSummary;
  }

  const props = defineProps<Props>();
  const value = useVModel(props, "modelValue");

  const items = ref<ItemSummary[]>([]);
  const form = ref({
    parent: null as ItemSummary | null,
    search: "",
  });

  // Whenever parent goes from value to null reset search
  watch(
    () => value.value,
    () => {
      if (!value.value) {
        form.value.search = "";
      }
    }
  );
</script>
