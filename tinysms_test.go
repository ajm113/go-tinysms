package tinysms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = &Options{
	Addr:     "smtp.gmail.com:587",
	Username: "tester@gmail.com",
	Password: "testing",
}

func TestTinySMS(t *testing.T) {
	c := NewClient(options)

	assert.NotNil(t, c)
	assert.NotNil(t, c.Options)
}
