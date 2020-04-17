package tinysms

import (
	"net/smtp"
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
