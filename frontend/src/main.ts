import { createApp } from "vue";

import { VueQueryPlugin } from "@tanstack/vue-query";

import App from "./App.vue";

import { router } from "./app/router";

import { queryClient } from "./app/providers/query-client";

import { setupInterceptors } from "./shared/api/interceptors";

import { bootstrap } from "@/modules/auth/composables/useAuth.ts" // tu composable bootstrap

import { setupAuthGuard } from "@/modules/auth/guards/auth.guard.ts"

import { createPinia } from "pinia"

import "@/style.css"

setupInterceptors();

const app = createApp(App);

const pinia = createPinia()

app.use(pinia)
app.use(router);
app.use(VueQueryPlugin, { queryClient });

await bootstrap()

setupAuthGuard(router)

app.mount("#app");