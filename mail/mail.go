package mail

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Client for sending mail via our mail service.
type Client struct {
	// client is the underlying type we make use of
	client *sendgrid.Client

	// Email is the sender for emails sent by this Client.
	Email *mail.Email
}

// NewClient creates a new mail client.
func NewClient(email *mail.Email) (client *Client) {
	client = &Client{}
	apiKey := os.Getenv("SENDGRID_API_KEY")
	client.client = sendgrid.NewSendClient(apiKey)
	client.Email = email
	return client
}

// Send an email with the provided details.
func (c Client) Send(s Message) (err error) {
	msg := mail.NewSingleEmail(c.Email, s.Subject, s.To, s.Text, s.HTML)
	msg.ReplyTo = s.ReplyTo
	_, err = c.client.Send(msg)
	return err
}

// Message stores the contents of an email to send using a Client.
type Message struct {
	// To is the recipient for the email.
	To *mail.Email

	// Subject is the subject for the email.
	Subject string

	// Text is the plaintext body for the email.
	Text string

	// HTML is the HTML body for the email.
	HTML string

	// ReplyTo optionally sets the address to which replies should be sent.
	ReplyTo *mail.Email
}

// NewEmail is a wrapper creating a new Email the client can use.
func NewEmail(name, address string) *mail.Email {
	return mail.NewEmail(name, address)
}
