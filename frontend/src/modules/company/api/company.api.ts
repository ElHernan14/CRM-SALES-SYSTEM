import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import {
  CompanyMeSchema,
  UploadCompanyCoverResponseSchema,
  UploadCompanyLogoResponseSchema,
  type CompanyMe,
  type UpdateCompanyMeRequest,
  type UploadCompanyCoverResponse,
  type UploadCompanyLogoResponse,
} from '../types/company.types';

export async function getCompanyMe(): Promise<CompanyMe> {
  const response = await http.get('/company/me');

  return unwrapResponse(CompanyMeSchema, response.data);
}

export async function updateCompanyMe(payload: UpdateCompanyMeRequest): Promise<CompanyMe> {
  const response = await http.patch('/company/me', payload);

  return unwrapResponse(CompanyMeSchema, response.data);
}

export async function uploadCompanyLogo(image: File): Promise<UploadCompanyLogoResponse> {
  const formData = new FormData();

  formData.append('logo', image);

  const response = await http.post('/company/me/logo', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });

  return unwrapResponse(UploadCompanyLogoResponseSchema, response.data);
}

export async function uploadCompanyCover(image: File): Promise<UploadCompanyCoverResponse> {
  const formData = new FormData();

  formData.append('cover_image', image);

  const response = await http.post('/company/me/cover-image', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });

  return unwrapResponse(UploadCompanyCoverResponseSchema, response.data);
}
