package tradier

import (
	"strconv"
	"strings"
	"time"
)

// Extract quota violation expiration from body message.
func parseQuotaViolationExpiration(body string) time.Time {
	if !strings.HasPrefix(body, "Quota Violation") {
		return time.Time{}
	}

	parts := strings.Fields(body)
	ms, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(ms/1000, 0)
}
