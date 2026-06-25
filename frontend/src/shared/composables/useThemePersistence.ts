import { watch } from "vue"
import { useUiStore } from "@/shared/stores/ui.store"

function applyTheme(theme: "light" | "dark") {
  const root = document.documentElement

  root.classList.remove("light", "dark")
  root.classList.add(theme)
}

export function initializeTheme() {
  const ui = useUiStore()

  const savedTheme = localStorage.getItem("theme")

  if (savedTheme === "light" || savedTheme === "dark") {
    ui.setTheme(savedTheme)
  } else {
    ui.setTheme("light")
  }

  applyTheme(ui.theme)

  watch(
    () => ui.theme,
    (theme) => {
      localStorage.setItem("theme", theme)
      applyTheme(theme)
    },
    {
      immediate: true
    }
  )
}