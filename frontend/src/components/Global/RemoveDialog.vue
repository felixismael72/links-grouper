<template>
  <div>
    <q-dialog v-model="open" persistent>
      <q-card>
        <q-card-section class="row items-center">
          <q-avatar icon="delete" color="secondary" text-color="white" />
          <span class="q-ml-sm">{{ removeMessage }}</span>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn outline class="text-capitalize" label="Cancelar" color="secondary" @click="$emit('closeDialog')" />
          <q-btn class="text-capitalize" label="Remover" color="secondary" @click="remove(), $emit('closeDialog')" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script>
import { useStore } from 'src/store';
import { computed } from 'vue';
export default {
  name: 'RemoveDialogComponent',
  props: ['dialogState', 'store', 'removeAction', 'removeMessage'],
  emits: ['closeDialog'],
  setup(props) {
    const $store = useStore();
    let open = computed(() => props.dialogState);
    let remove = () => {
      $store.dispatch(props.store + '/' + props.removeAction);
    };
    return {
      open,
      remove,
    };
  },
};
</script>
