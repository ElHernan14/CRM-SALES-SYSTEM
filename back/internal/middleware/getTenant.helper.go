package middleware

import "context"

func GetTenant(ctx context.Context) TenantContext {
	tenant, ok := ctx.Value(TenantContextKey).(TenantContext)
	if !ok {
		return TenantContext{}
	}
	return tenant
}
