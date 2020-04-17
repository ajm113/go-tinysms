package tinysms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = &Options{
	Addr:        "smtp.gmail.com:587",
	Username:    "tester@gmail.com",
	Password:    "testing",
	FromAddress: "andrew@test.com",
}

func TestTinySMS(t *testing.T) {
	c := NewClient(options)

	assert.NotNil(t, c)
	assert.NotNil(t, c.Options)
}

func TestSend(t *testing.T) {
	c := NewClient(options)

	err := c.Send("5555555555", "AT&T", "tinysms Golang Library is 1337!")

	assert.NotNil(t, err)
}

func TestSendErrors(t *testing.T) {
	c := NewClient(options)

	err := c.Send("5555555555", "error", "failed message!")

	assert.NotNil(t, err)

	c.Options.Addr = ":123"
	err = c.Send("5555555555", "AT&T", "failed message!")

	assert.NotNil(t, err)
}
