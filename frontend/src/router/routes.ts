import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/auth',
    component: () => import('layouts/RedirectLayout.vue'),
    children: [
      {
        path: '/auth/signIn',
        component: () => import('pages/SignInPage.vue')
      },
      {
        path: '/auth/signUp',
        component: () => import('pages/SignUpPage.vue')
      },
    ],
  },
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '/',
        beforeEnter: (to, from, next) => {
          if (localStorage.getItem('token')) {
            next();
          } else {
            next('/auth/signIn');
          }
        },
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '/groups',
        component: () => import('pages/GroupsLinks.vue')
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
