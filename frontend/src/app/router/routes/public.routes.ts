import PublicLayout from '@/app/layouts/PublicLayout.vue';
import LoginPage from '@/modules/auth/pages/LoginPage.vue';

export const publicRoutes = [
  {
    path: '/',
    component: PublicLayout,
    children: [
      {
        path: '',
        redirect: { name: 'login' },
      },
      {
        path: 'login',
        name: 'login',
        component: LoginPage,
        meta: {
          guestOnly: true,
        },
      },
    ],
  },
];
