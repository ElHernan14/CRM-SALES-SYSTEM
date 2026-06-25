import { computed } from "vue"
import { useUiStore } from "@/shared/stores/ui.store"

export function useTheme() {
  const ui = useUiStore()

  const theme = computed(() => ui.theme)
  const isDark = computed(() => ui.isDark)

  function setTheme(theme: "light" | "dark") {
    ui.setTheme(theme)
  }

  function toggleTheme() {
    ui.toggleTheme()
  }

  return {
    theme,
    isDark,
    setTheme,
    toggleTheme
  }
}