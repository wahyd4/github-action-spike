package spike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func A_Test(t *testing.T) {
	a := "he2llo"
	if a != hello() {
		t.Errorf("a %v is not what we expected %v ", a, "hello")
	}
}

func Test(t *testing.T) {
	assert.Nil(t, 123, "It Should not be nil")
}
