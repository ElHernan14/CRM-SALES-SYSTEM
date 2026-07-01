import { createRouter, createWebHistory } from "vue-router";
import PublicLayout from "@/app/layouts/PublicLayout.vue";
import ERPLayout from "@/modules/erp/layouts/ERPLayout.vue"
import DashboardPage from "@/modules/erp/pages/DashboardPage.vue"
import LoginPage from "@/modules/auth/pages/LoginPage.vue";

const routes = [
  {
    path: "/",
    component: PublicLayout,
    children: [
      // {
      //   path: "",
      //   component: LandingPage
      // },
      {
        path: "",
        redirect: { name: "login" }
      },
      {
        path: "login",
        name:"login",
        component: LoginPage,
        meta: {
          guestOnly: true
        }
      }
    ]
  },
  {
    path: "/erp",
    component: ERPLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: "dashboard",
        name: "erp-dashboard",
        component: DashboardPage,
      },
    ],
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
