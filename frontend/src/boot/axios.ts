import { boot } from 'quasar/wrappers'
import axios, { AxiosInstance } from 'axios'
import { getLocalToken } from 'src/store/utils';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}

const api = axios.create({ baseURL: 'http://localhost:8000/api/' })
const auth = axios.create({ baseURL: 'http://localhost:8001/api/auth/' })
            
export default boot(({ router, store, app }) => {
  app.config.globalProperties.$axios = axios

  app.config.globalProperties.$api = api
  app.config.globalProperties.$auth = auth

  const token = getLocalToken();
  if (token && !store.getters['user/isAuthenticated']) {
    store.commit('user/setAuth', token);
  }

  setInterval(() => {
    if (store.getters['user/isAuthenticated']) {
      store
        .dispatch('user/refreshUser')
        .then(() => {
          console.log('refreshed');
        })
        .catch(() => {
          router.push('/');
          store.commit('user/unsetAuth');
        });
    }
  }, 5000);
});

export { api, auth }
