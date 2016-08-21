package less

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestFloat64WithTie(t *testing.T) {
	trueFCalled := false
	trueF := func() bool {
		trueFCalled = true
		return true
	}
	assert.True(t, "1 < 2", Float64WithTie(1., 2., trueF))
	assert.False(t, "trueFCalled", trueFCalled)

	assert.False(t, "2 < 1", Float64WithTie(2., 1., trueF))
	assert.False(t, "trueFCalled", trueFCalled)

	assert.True(t, "1 < 1 with true tie", Float64WithTie(1., 1., trueF))
	assert.True(t, "trueFCalled", trueFCalled)
}

func TestIntWithTie(t *testing.T) {
	trueFCalled := false
	trueF := func() bool {
		trueFCalled = true
		return true
	}
	assert.True(t, "1 < 2", IntWithTie(1, 2, trueF))
	assert.False(t, "trueFCalled", trueFCalled)

	assert.False(t, "2 < 1", IntWithTie(2, 1, trueF))
	assert.False(t, "IntWithTie", trueFCalled)

	assert.True(t, "1 < 1 with true tie", IntWithTie(1, 1, trueF))
	assert.True(t, "trueFCalled", trueFCalled)
}

func TestStrLenOrContent(t *testing.T) {
	assert.True(t, "2 < 11", StrLenOrContent("2", "11"))
	assert.False(t, "11 < 2", StrLenOrContent("11", "2"))
	assert.True(t, "1 < 22", StrLenOrContent("1", "22"))
	assert.True(t, "1 < 2", StrLenOrContent("1", "2"))
}
