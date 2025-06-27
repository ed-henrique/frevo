package duration

import (
	"time"

	"github.com/ed-henrique/frevo/internal/assert"
	"github.com/ed-henrique/frevo/internal/types"
)

const (
	errNegativeValue = "the value passed is negative."
)

func Nanoseconds[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Nanosecond
}

func Microseconds[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Microsecond
}

func Milliseconds[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Millisecond
}

func Seconds[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Second
}

func Minutes[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Minute
}

func Hours[T types.Number](x T) time.Duration {
	assert.AssertTrue(x >= 0, errNegativeValue)
	return time.Duration(x) * time.Hour
}
