package tinysms

import (
	"net/smtp"
)

type (
	// SMSClient Client object used to send SMS messages.
	SMSClient struct {
		options    *Options
		stmpClient *smtp.Client
	}
)
