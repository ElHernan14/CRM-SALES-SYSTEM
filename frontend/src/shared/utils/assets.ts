import { env } from '@/shared/config/env';

export function getAssetUrl(path?: string | null) {
  if (!path) return null;

  if (path.startsWith('http')) {
    return path;
  }

  return `${env.assetsUrl}${path}`;
}
