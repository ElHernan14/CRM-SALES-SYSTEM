import { http } from '@/shared/api/http';

import { unwrapResponse } from '@/shared/utils/response';

import { DashboardOverviewSchema, type DashboardOverview } from '../types/dashboard.types';

export async function getDashboardOverview(): Promise<DashboardOverview> {
  const response = await http.get('/dashboard/overview');

  return unwrapResponse(DashboardOverviewSchema, response.data);
}
