const apiUrl = import.meta.env.VITE_API_URL;

if (!apiUrl) {
  throw new Error('Missing VITE_API_URL environment variable.');
}

export const env = {
  apiUrl,
  assetsUrl: import.meta.env.VITE_ASSETS_URL || apiUrl.replace(/\/api\/?$/, ''),
} as const;
