# tinysms
**Simple free SMS client using SMTP**

[![codecov](https://codecov.io/gh/ajm113/go-tinysms/branch/master/graph/badge.svg)](https://codecov.io/gh/ajm113/go-tinysms)
[![Build Status](https://travis-ci.org/ajm113/go-tinysms.svg?branch=master)](https://travis-ci.org/ajm113/go-tinysms)
[![Go Report Card](https://goreportcard.com/badge/github.com/ajm113/go-tinysms)](https://goreportcard.com/report/github.com/ajm113/go-tinysms)

This is a library is perfect for small projects you are only sending a small amount of SMS to a single phone or a few you know of. Unfortunately, you need to know the user's carrier to send SMS. If you don't want to have to ask your users that. Then this project is not for you. I recommend using [Twilio](https://www.twilio.com/) for this if you plan on high volumn SMS messages. 

## Requirements

You are required to have an SMTP server to send SMS. Google provides this for free including few other major email providers. You will need to look up or ask your provider to give you SMTP access to send messages. **This library currently only supports major US wireless carriers**. If you would like to add one. Please make a PR request or open an issue!

## Install

`go get github.com/ajm113/go-tinysms`

## Usuage

```
import "github.com/ajm113/go-tinysms"

func main() {
	c := tinysms.NewClient(&tinysms.Options{
		// SMTP Server and login settings.
		Addr:        "smtp.gmail.com:587",
		Username:    "tester@gmail.com",
		Password:    "testing",

		// How you want the from address to show on text messages.
		FromAddress: "andrew@test.com",
	})

	// Send a text message to a phone using AT&T.
	err := c.Send("5555555555", "AT&T", "tinysms Golang library is 1337!")
}
```

## This project wants help!
Few things you can help improve the project!

- Better carrier support of international carriers.
- Improved documentation.
- Better code coverage and unit tests.

## License

MIT

## Authors

- Andrew McRobb
