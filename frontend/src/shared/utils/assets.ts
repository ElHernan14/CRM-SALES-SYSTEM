import { env } from '@/shared/config/env';

export function getPublicAssetUrl(path?: string | null): string | null {
  if (!path) return null;

  if (path.startsWith('http://') || path.startsWith('https://')) {
    return path;
  }

  const normalizedPath = path.startsWith('/') ? path : `/${path}`;

  return `${env.assetsUrl}${normalizedPath}`;
}

export function getCompanyLogoUrl(path?: string | null): string | null {
  if (!path) return null;

  if (path.includes('/')) {
    return getPublicAssetUrl(path);
  }

  return getPublicAssetUrl(`/uploads/company/logos/${path}`);
}

export function getCompanyCoverUrl(path?: string | null): string | null {
  if (!path) return null;

  if (path.includes('/')) {
    return getPublicAssetUrl(path);
  }

  return getPublicAssetUrl(`/uploads/company/covers/${path}`);
}

export function getProductImageUrl(path?: string | null): string | null {
  if (!path) return null;

  if (path.includes('/')) {
    return getPublicAssetUrl(path);
  }

  return getPublicAssetUrl(`/uploads/product/images/${path}`);
}
