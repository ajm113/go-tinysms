package tinysms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTinySMS(t *testing.T) {
	e := NewClient(&Options{
		Addr:     "localhost:1234",
		Username: "test@email.com",
		Password: "password",
	})

	assert.NotNil(t, e)
}
