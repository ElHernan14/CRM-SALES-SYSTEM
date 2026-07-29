import { computed, ref } from 'vue';

import { brand } from '@/shared/config/brand';

function updateBrowserThemeColor(isDark: boolean) {
  const themeColorMeta = document.querySelector<HTMLMetaElement>('meta[name="theme-color"]');

  if (!themeColorMeta) return;

  themeColorMeta.content = isDark
    ? brand.metadata.themeColor.dark
    : brand.metadata.themeColor.light;
}

export type ThemePreference = 'light' | 'dark' | 'system';

const STORAGE_KEY = 'nexora-theme';

const theme = ref<ThemePreference>('system');
const systemIsDark = ref(false);

let initialized = false;
let systemMediaQuery: MediaQueryList | null = null;

function getStoredTheme(): ThemePreference {
  const storedTheme = localStorage.getItem(STORAGE_KEY);

  if (storedTheme === 'light' || storedTheme === 'dark' || storedTheme === 'system') {
    return storedTheme;
  }

  return 'system';
}

function resolveIsDark(preference: ThemePreference) {
  if (preference === 'dark') {
    return true;
  }

  if (preference === 'light') {
    return false;
  }

  return systemIsDark.value;
}

function applyTheme(preference: ThemePreference) {
  const dark = resolveIsDark(preference);
  const root = document.documentElement;

  root.classList.toggle('dark', dark);
  updateBrowserThemeColor(dark);

  root.style.colorScheme = dark ? 'dark' : 'light';
}

function handleSystemThemeChange(event: MediaQueryListEvent) {
  systemIsDark.value = event.matches;

  if (theme.value === 'system') {
    applyTheme('system');
  }
}

function initializeTheme() {
  if (initialized || typeof window === 'undefined') {
    return;
  }

  systemMediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
  systemIsDark.value = systemMediaQuery.matches;

  theme.value = getStoredTheme();

  applyTheme(theme.value);

  systemMediaQuery.addEventListener('change', handleSystemThemeChange);

  initialized = true;
}

export function useTheme() {
  initializeTheme();

  const isDark = computed(() => {
    return resolveIsDark(theme.value);
  });

  function setTheme(preference: ThemePreference) {
    theme.value = preference;

    localStorage.setItem(STORAGE_KEY, preference);

    applyTheme(preference);
  }

  function toggleTheme() {
    setTheme(isDark.value ? 'light' : 'dark');
  }

  return {
    theme,
    isDark,
    setTheme,
    toggleTheme,
  };
}
