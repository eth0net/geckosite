package mail

import (
	"os"
	"sync"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	once   sync.Once
	client *sendgrid.Client
)

// Init sets up the mail client for use.
func Init() {
	once.Do(func() {
		client = sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	})
}

// NewEmail wrapper.
func NewEmail(name, address string) *mail.Email {
	return mail.NewEmail(name, address)
}

// Send an email.
func Send(to *mail.Email, subject, bodyPlainText, bodyHTML string) error {
	from := NewEmail("Enquiries", "enquiries@hollyshatchlings.co.uk")
	message := mail.NewSingleEmail(from, subject, to, bodyPlainText, bodyHTML)

	_, err := client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
