import { defineStore } from "pinia"

export type ThemeMode = "light" | "dark"

interface UiState {
  loading: boolean
  loadingMessage: string | null

  theme: ThemeMode
}

export const useUiStore = defineStore("ui", {
  state: (): UiState => ({
    loading: false,
    loadingMessage: null,

    theme: "light"
  }),

  getters: {
    isDark: (state) => state.theme === "dark"
  },

  actions: {
    setLoading(
      loading: boolean,
      message?: string
    ) {
      this.loading = loading
      this.loadingMessage = loading
        ? message ?? null
        : null
    },

    setTheme(theme: ThemeMode) {
      this.theme = theme
    },

    toggleTheme() {
      this.theme =
        this.theme === "light"
          ? "dark"
          : "light"
    }
  }
})