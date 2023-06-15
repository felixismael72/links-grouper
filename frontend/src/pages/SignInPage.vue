<template>
  <q-page padding>
    <signIn class="fixed-center" />
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import signIn from 'src/components/Authentication/SignIn.vue';
import { useStore } from 'src/store';
export default defineComponent({
  name: 'SignInPage',
  components: { signIn },
  setup() {
    const $store = useStore();

    const pathData = {
      paths: ['Início', 'Entrar'],
      icons: ['home', 'sigin'],
    };

    $store.commit('path/setPath', pathData);
  },
});
</script>
