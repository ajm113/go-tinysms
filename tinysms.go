package github.com/ajm113/go-tinysms


import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"net/smtp"
	"strings"
)

type (
	SMSClient struct {
		options *Options
		stmpClient *smtp.Client,
	}
)
