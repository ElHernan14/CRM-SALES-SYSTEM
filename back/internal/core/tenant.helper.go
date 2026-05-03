package core

import (
	"context"
	"log"
)

type TenantContext struct {
	UserID      int
	Email       string
	CompanyID   *int
	Roles       []string
	Permissions []string
}

type contextKey string

const TenantContextKey contextKey = "tenant"

func GetTenant(ctx context.Context) *TenantContext {
	tenant, ok := ctx.Value(TenantContextKey).(*TenantContext)
	if !ok || tenant == nil {
		log.Printf("No tenant found in context")
		return nil
	}
	return tenant
}
