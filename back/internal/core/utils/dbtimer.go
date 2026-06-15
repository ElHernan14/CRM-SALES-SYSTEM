package utils

import (
	"context"
	"log"
	"time"
)

const SlowQueryThreshold = 200 * time.Millisecond

func LogSlowQuery(
	ctx context.Context,
	operation string,
	table string,
	start time.Time,
	rowsAffected int64,
) {

	duration := time.Since(start)

	if duration < SlowQueryThreshold {
		return
	}

	requestID := GetRequestID(ctx)

	log.Printf(
		"DB slow query | request_id=%s | operation=%s | table=%s | duration=%s | rowsAffected=%d",
		requestID,
		operation,
		table,
		duration,
		rowsAffected,
	)
}
