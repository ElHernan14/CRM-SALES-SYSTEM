import { http } from '@/shared/api/http';
import { unwrapResponse } from '@/shared/utils/response';
import {
  // ProductTypeSchema
  ProductDetailResponseSchema,
  type ProductDetailResponse,
  type UpdateProductRequest,
  // ProductListItemSchema
  GetProductsResponseSchema,
  type GetProductsRequest,
  type GetProductsResponse,
  // CreateProductSchema
  type CreateProductRequest,
  // Image
  UploadProductImageResponseSchema,
  type UploadProductImageResponse,
} from '../types/product.types';

export async function getProducts(params: GetProductsRequest): Promise<GetProductsResponse> {
  const response = await http.get('/product', {
    params,
  });

  return unwrapResponse(GetProductsResponseSchema, response.data);
}

export async function createProduct(payload: CreateProductRequest): Promise<ProductDetailResponse> {
  const response = await http.post('/product', payload);

  return unwrapResponse(ProductDetailResponseSchema, response.data);
}

export async function updateProduct(
  id: number,
  payload: UpdateProductRequest
): Promise<ProductDetailResponse> {
  const response = await http.patch(`/product/${id}`, payload);

  return unwrapResponse(ProductDetailResponseSchema, response.data);
}

export async function deleteProduct(id: number): Promise<void> {
  await http.delete(`/product/${id}`);
}

export async function uploadProductImage(
  productId: number,
  image: File
): Promise<UploadProductImageResponse> {
  const formData = new FormData();

  formData.append('image', image);

  const response = await http.post(`/product/${productId}/image`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });

  return unwrapResponse(UploadProductImageResponseSchema, response.data);
}
