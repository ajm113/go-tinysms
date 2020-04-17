package tinysms

// Options SMTP Connection strings.
type Options struct {
	// Addr we are connecting to. e.g: "smtp.gmail.com:587"
	Addr string

	// Username Sometimes email address or username used to login
	Username string
	Password string

	// FromAddress Use custom from address if service supports it. Sometimes they will ignore this.
	FromAddress string
}
