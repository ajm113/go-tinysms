package tinysms

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

type (
	// SMSClient Client object used to send SMS messages.
	SMSClient struct {
		Options    *Options
		stmpClient *smtp.Client
	}
)

// NewClient Creates new client instant to send SMS.
func NewClient(options *Options) (c *SMSClient) {
	c = &SMSClient{Options: options}
	return
}

// Send Sends SMS message to target number supplied with message string.
func (c *SMSClient) Send(number, carrier, message string) (err error) {
	carrierEmailDomain, ok := Carriers[strings.ToLower(carrier)]

	if !ok {
		err = errors.New("unsupported carrier " + carrier + "! Please check available carriers list!")
		return
	}

	to := []string{fmt.Sprintf("%s@%s", number, carrierEmailDomain)}

	// Now split the host/port for smtp.PlainAuth
	host, _, err := net.SplitHostPort(c.Options.Addr)

	if err != nil {
		return
	}

	auth := smtp.PlainAuth("", c.Options.Username, c.Options.Password, host)

	// Now send the request!
	err = smtp.SendMail(c.Options.Addr, auth, c.Options.FromAddress, to, []byte(message))

	return
}
