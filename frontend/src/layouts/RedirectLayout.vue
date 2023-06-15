<template>
  <q-layout view="hHh Lpr fFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title> Links Grouper </q-toolbar-title>
        <div><q-btn color="white" text-color="secondary" class="text-capitalize" icon="home" to="/" label="Início" />
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  name: 'AuthenticationLayout',
  setup() {
    const leftDrawer = ref(false);
    return { leftDrawer };
  },
});
</script>
