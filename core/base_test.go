package core

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	i := Foo()
	assert.Equal(t, i, "bar")
}

func TestRandom(t *testing.T) {
	i := Random()
	s, err := strconv.ParseFloat(i, 64)
	if err != nil {
		t.Error("can't parse result")
	}
	assert.GreaterOrEqual(t, s, 0.0)
	assert.LessOrEqual(t, s, 1.0)
}

func TestWhoIS(t *testing.T) {
	sampleIP := [3]string{"8.8.8.8", "1.2.3.4", "79.176.20.131"}

	for _, ip := range sampleIP {
		i := WhoIS(ip)
		assert.NotEqual(t, i, "oops")
	}

}
