package tradier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_parseQuotaViolationExpiration(t *testing.T) {
	t.Run("Missing quota prefix", func(t *testing.T) {
		output := parseQuotaViolationExpiration("")
		assert.Equal(t, output.Unix(), time.Time{}.Unix())
	})

	t.Run("Quota Violation not a number", func(t *testing.T) {
		output := parseQuotaViolationExpiration("Quota Violation not a number")
		assert.Equal(t, output.Unix(), time.Time{}.Unix())
	})

	t.Run("Quota Violation is valid", func(t *testing.T) {
		expiration := time.Now().Add(time.Minute)

		output := parseQuotaViolationExpiration(fmt.Sprintf("Quota Violation expires in %v000", expiration.Unix()))
		assert.Equal(t, output.Unix(), expiration.Unix())
	})
}
