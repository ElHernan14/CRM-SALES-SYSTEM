import { createRouter, createWebHistory } from "vue-router";
import PublicLayout from "@/app/layouts/PublicLayout.vue";
import ERPLayout from "@/app/layouts/ERPLayout.vue";
import LoginPage from "@/modules/auth/pages/LoginPage.vue";
import DashboardPage from "@/modules/dashboard/pages/DashboardPage.vue";
import ProductsPage from "@/modules/products/pages/ProductPage.vue";

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
        path: "login",
        name:"login",
        component: LoginPage
      }
    ]
  },
  {
    path: "/erp",
    component: ERPLayout,
    children: [
      {
        path: "dashboard",
        component: DashboardPage
      },
      {
        path: "products",
        component: ProductsPage
      }
    ]
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
