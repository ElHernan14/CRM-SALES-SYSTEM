import { createApp } from "vue";

import { VueQueryPlugin } from "@tanstack/vue-query";

import App from "./App.vue";

import { router } from "./app/router";

import { queryClient } from "./app/providers/query-client";

import { setupInterceptors } from "./shared/api/interceptors";

import { bootstrap } from "@/modules/auth/composables/useAuth.ts" // tu composable bootstrap

setupInterceptors();

const app = createApp(App);

app.use(router);

app.use(VueQueryPlugin, {
  queryClient,
});

await bootstrap

app.mount("#app");