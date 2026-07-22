import type { RouteRecordRaw } from 'vue-router';

import PublicLayout from '@/app/layouts/PublicLayout.vue';

const LandingPage = () => import('@/modules/public/pages/LandingPage.vue');

const LoginPage = () => import('@/modules/auth/pages/LoginPage.vue');

const RegisterPage = () => import('@/modules/auth/pages/RegisterPage.vue');

const PersonalRegisterPage = () => import('@/modules/auth/pages/PersonalRegisterPage.vue');

const BusinessRegisterPage = () => import('@/modules/auth/pages/BusinessRegisterPage.vue');

const InvitationRegisterPage = () => import('@/modules/auth/pages/InvitationRegisterPage.vue');

export const publicRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    component: PublicLayout,

    children: [
      {
        path: '',
        name: 'landing',
        component: LandingPage,
      },

      {
        path: 'login',
        name: 'login',
        component: LoginPage,

        meta: {
          guestOnly: true,
        },
      },

      {
        path: 'register',
        name: 'register',
        component: RegisterPage,

        meta: {
          guestOnly: true,
        },
      },

      {
        path: 'register/personal',
        name: 'register-personal',
        component: PersonalRegisterPage,

        meta: {
          guestOnly: true,
        },
      },

      {
        path: 'register/business',
        name: 'register-business',
        component: BusinessRegisterPage,

        meta: {
          guestOnly: true,
        },
      },

      {
        path: 'register/invitation/:token',
        name: 'register-invitation',
        component: InvitationRegisterPage,

        props: true,

        meta: {
          guestOnly: true,
        },
      },
    ],
  },
];
