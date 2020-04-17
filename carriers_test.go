package tinysms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarriers(t *testing.T) {
	assert.NotNil(t, Carriers)
	assert.NotEmpty(t, Carriers)
}
