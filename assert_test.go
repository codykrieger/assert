package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatItPanicsOnFalseConditions(t *testing.T) {
	assert.Panics(t, func() {
		Assert(false)
	}, "Calling Assert() with a false condition should panic.")
}

func TestThatItDoesNotPanicOnTrueConditions(t *testing.T) {
	assert.NotPanics(t, func() {
		Assert(true)
	}, "Calling Assert() with a true condition should not panic.")
}
