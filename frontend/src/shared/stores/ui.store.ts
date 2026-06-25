import { defineStore } from "pinia"

export type ThemeMode = "light" | "dark"

interface ConfirmDialogOptions {
  title: string
  description?: string
  confirmText?: string
  cancelText?: string
  variant?: "default" | "destructive"
  onConfirm?: () => void | Promise<void>
}

interface UiState {
  loading: boolean
  loadingMessage: string | null
  theme: ThemeMode

  confirmOpen: boolean
  confirmTitle: string
  confirmDescription: string | null
  confirmText: string
  cancelText: string
  confirmVariant: "default" | "destructive"
  confirmAction: (() => void | Promise<void>) | null
}

export const useUiStore = defineStore("ui", {
  state: (): UiState => ({
    loading: false,
    loadingMessage: null,
    theme: "light",
    
    confirmOpen: false,
    confirmTitle: "",
    confirmDescription: null,
    confirmText: "Confirm",
    cancelText: "Cancel",
    confirmVariant: "default",
    confirmAction: null,
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
    },

    openConfirm(options: ConfirmDialogOptions) {
      this.confirmOpen = true
      this.confirmTitle = options.title
      this.confirmDescription = options.description ?? null
      this.confirmText = options.confirmText ?? "Confirm"
      this.cancelText = options.cancelText ?? "Cancel"
      this.confirmVariant = options.variant ?? "default"
      this.confirmAction = options.onConfirm ?? null
    },

    closeConfirm() {
      this.confirmOpen = false
      this.confirmTitle = ""
      this.confirmDescription = null
      this.confirmText = "Confirm"
      this.cancelText = "Cancel"
      this.confirmVariant = "default"
      this.confirmAction = null
    },

    async confirm() {
      if (this.confirmAction) {
        await this.confirmAction()
      }

      this.closeConfirm()
    }
  }
})