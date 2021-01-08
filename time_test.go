package tradier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTime_Set(t *testing.T) {
	t.Run("Parse 2006-01-02T15:04:05", func(t *testing.T) {
		input := "2006-01-02T15:04:05"
		value := DateTime{}
		err := value.Set(input)
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(1136214245))
	})

	t.Run("Parse 2006-01-02", func(t *testing.T) {
		input := "2006-01-02"
		value := DateTime{}
		err := value.Set(input)
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(1136160000))
	})

	t.Run("Parse 15:04", func(t *testing.T) {
		input := "15:04"
		value := DateTime{}
		err := value.Set(input)
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(-62167164960))
	})

	t.Run("Parse seconds since epoc", func(t *testing.T) {
		input := time.Now()
		ms := fmt.Sprintf("%v%v", input.Unix(), input.Nanosecond()/1000000)

		value := DateTime{}
		err := value.Set(ms)
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), input.Unix())
	})

	t.Run("Invalid time", func(t *testing.T) {
		input := "not a time"

		value := DateTime{}
		err := value.Set(input)
		assert.Error(t, err)
	})
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	t.Run("Parse 2006-01-02T15:04:05", func(t *testing.T) {
		input := "2006-01-02T15:04:05"
		value := DateTime{}
		err := value.UnmarshalJSON([]byte(input))
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(1136214245))
	})

	t.Run("Parse 2006-01-02", func(t *testing.T) {
		input := "2006-01-02"
		value := DateTime{}
		err := value.UnmarshalJSON([]byte(input))
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(1136160000))
	})

	t.Run("Parse 15:04", func(t *testing.T) {
		input := "15:04"
		value := DateTime{}
		err := value.UnmarshalJSON([]byte(input))
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), int64(-62167164960))
	})

	t.Run("Parse seconds since epoc", func(t *testing.T) {
		input := time.Now()
		ms := fmt.Sprintf("%v000", input.Unix())

		value := DateTime{}
		err := value.UnmarshalJSON([]byte(ms))
		assert.NoError(t, err)
		assert.Equal(t, value.Unix(), input.Unix())
	})

	t.Run("Invalid time", func(t *testing.T) {
		input := "not a time"

		value := DateTime{}
		err := value.UnmarshalJSON([]byte(input))
		assert.Error(t, err)
	})
}

func TestParseTimeMs(t *testing.T) {
	t.Run("Invalid ms", func(t *testing.T) {
		_, err := ParseTimeMs("not a number")
		assert.Error(t, err)
	})

	t.Run("Invalid ms", func(t *testing.T) {
		output, err := ParseTimeMs("123456")
		assert.NoError(t, err)
		assert.Equal(t, output.Nanosecond(), 456000000)
	})
}
