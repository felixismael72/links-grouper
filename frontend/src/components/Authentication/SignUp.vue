<template>
  <div class="q-pa-md" style="max-width: 400px">
    <q-form @submit="save()" class="q-gutter-md">
      <q-input filled v-model="username" label="Username" lazy-rules :rules="[
        (val) => (val && val.length > 0) || 'Este campo é obrigatório',
      ]" />
      
      <q-input filled v-model="full_name" label="Full Name" lazy-rules :rules="[
        (val) => (val && val.length > 0) || 'Este campo é obrigatório',
      ]" />

      <q-input filled v-model="email" label="E-mail" lazy-rules :rules="[
        (val) => (val && val.length > 0) || 'Este campo é obrigatório',
      ]" />

      <q-input v-model="password" filled label="Senha" :type="isPwd ? 'password' : 'text'" lazy-rules :rules="[
        (val) => (val && val.length > 0) || 'Este campo é obrigatório',
      ]">
        <template v-slot:append>
          <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer" @click="isPwd = !isPwd" />
        </template>
      </q-input>

      <div v-ripple class="relative-position container flex flex-center text-white">
        <q-btn label="Cadastrar" type="submit" color="secondary" style="min-width: 150px" />
      </div>
    </q-form>
  </div>
</template>

<script lang="ts">
import { useStore } from 'src/store';
import { defineComponent, ref } from 'vue';
import { useRouter } from 'vue-router';
export default defineComponent({
  name: 'SignUpComponent',
  setup() {
    const $store = useStore();
    const $router = useRouter();

    const password = ref('');
    const isPwd = ref(true);
    const email = ref('');
    const username = ref('');
    const full_name = ref('')

    const save = () => {
      $store.dispatch('user/signUpUser', {
        username: username.value,
        full_name: full_name.value,
        email: email.value,
        password: password.value,
      });

      $router.push('/');
    };

    return {
      password,
      isPwd,
      email,
      username,
      full_name,
      save,
    };
  },
});
</script>
