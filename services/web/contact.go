package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/eth0net/geckosite/mail"
	"github.com/eth0net/geckosite/templates"
	"github.com/gin-gonic/gin"
)

const (
	basePlain = "Name: %v\r\nEmail: %v\r\nMessage: %v\r\n"
	baseHTML  = `<div><b style="color:orange">Name:</b> %v</div>
<div><b style="color:orange">Email:</b> %v</div>
<div><b style="color:orange">Message:</b> %v</div>`
)

// Contact returns the contact us page.
func (s service) Contact(c *gin.Context) {
	data := struct {
		Success                          bool
		Title, Path, Message, CaptchaKey string
	}{
		Title:      "Contact Us",
		Path:       c.Request.URL.Path,
		CaptchaKey: os.Getenv("HCAPTCHA_SITE_KEY"),
	}

	tmpl := templates.Parse("contact")
	defer tmpl.ExecuteTemplate(c.Writer, "layout", &data)

	if c.Request.Method == http.MethodPost {
		values := struct{ Name, Email, Message, Token string }{
			c.Request.FormValue("name"),
			c.Request.FormValue("email"),
			c.Request.FormValue("message"),
			c.Request.FormValue("h-captcha-response"),
		}

		res, err := http.PostForm("https://hcaptcha.com/siteverify", url.Values{
			"secret":   []string{os.Getenv("HCAPTCHA_SECRET_KEY")},
			"response": []string{values.Token},
			// "remoteip": []string{r.RemoteAddr},
			"sitekey": []string{os.Getenv("HCAPTCHA_SITE_KEY")},
		})
		if err != nil {
			log.Printf("Error occurred while verifying token with hcaptcha: %v\n", err)
			data.Message = "Failed to verify humanity!"
			return
		}

		bts, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error occurred while reading response from hcaptcha: %v\n", err)
			data.Message = "Failed to verify humanity!"
			return
		}

		var verification struct {
			Success   bool   `json:"success"`
			Timestamp string `json:"challenge_ts"`
			HostName  string `json:"hostname"`
			// Credit    bool     `json:"credit"`
			Errors []string `json:"error-codes"`
		}
		err = json.Unmarshal(bts, &verification)
		if err != nil {
			log.Printf("Error occurred while parsing JSON response: %v\n", err)
			data.Message = "Failed to verify humanity!"
			return
		}

		if !verification.Success {
			log.Printf("Verification failed with hcaptcha: %#v\n", verification)
			data.Message = "Failed to verify humanity!"
			return
		}

		email := mail.NewEmail("Enquiries", "enquiries@hollyshatchlings.co.uk")
		client := mail.NewClient(email)

		subject := "Website Enquiry from " + values.Name
		text := fmt.Sprintf(basePlain, values.Name, values.Email, values.Message)
		html := fmt.Sprintf(baseHTML, values.Name, values.Email, values.Message)

		replyTo := mail.NewEmail(values.Name, values.Email)

		msg := mail.Message{
			To:      email,
			Subject: subject,
			Text:    text,
			HTML:    html,
			ReplyTo: replyTo,
		}

		err = client.Send(msg)
		if err == nil {
			data.Success = true
		}
	}
}
