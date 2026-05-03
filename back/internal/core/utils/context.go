package utils

import "context"

type contextKey string

const RequestIDKey contextKey = "request_id"

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(RequestIDKey).(string); ok {
		return v
	}
	return ""
}
