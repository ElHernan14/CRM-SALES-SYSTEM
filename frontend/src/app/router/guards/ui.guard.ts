import type { Router } from 'vue-router';
import { useUiStore } from '@/shared/stores/ui.store';

export function setupUIGuard(router: Router) {
  router.beforeEach((to, from, next) => {
    const ui = useUiStore();

    if (to.path !== from.path) {
      ui.setLoading(true, 'Loading workspace...');
    }

    next();
  });

  router.afterEach(() => {
    const ui = useUiStore();

    setTimeout(() => {
      ui.setLoading(false);
    }, 250);
  });
}
