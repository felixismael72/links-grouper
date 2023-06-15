<template>
  <q-page padding>
    <SignUp class="fixed-center" />
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SignUp from 'components/Authentication/SignUp.vue';
import { useStore } from 'src/store';
export default defineComponent({
  name: 'SignUpPage',
  components: { SignUp },
  setup() {
    const $store = useStore();

    const pathData = {
      paths: ['Início', 'Cadastrar'],
      icons: ['home', 'app_registration'],
    };

    $store.commit('path/setPath', pathData);
  },
});
</script>
