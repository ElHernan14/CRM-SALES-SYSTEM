import { defineStore } from 'pinia';

export const useStoreUiStore = defineStore('store-ui', {
  state: () => ({
    cartOpen: false,
    mobileMenuOpen: false,
    searchOpen: false,
  }),

  actions: {
    openCart() {
      this.cartOpen = true;
    },

    closeCart() {
      this.cartOpen = false;
    },

    toggleCart() {
      this.cartOpen = !this.cartOpen;
    },

    setMobileMenuOpen(open: boolean) {
      this.mobileMenuOpen = open;
    },

    setSearchOpen(open: boolean) {
      this.searchOpen = open;
    },
  },
});
