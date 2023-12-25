package utils

import "context"

func GetTraceID(c context.Context) string {
	traceID := c.Value("X-Txcn-Trace-Id")
	if v, ok := traceID.(string); ok {
		return v
	}
	if v, ok := traceID.(*string); ok {
		return *v
	}
	return ""
}
