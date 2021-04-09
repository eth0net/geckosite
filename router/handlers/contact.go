package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/raziel2244/geckosite/mail"
)

// Contact returns the contact us page.
func Contact(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Success                          bool
		Title, Path, Message, CaptchaKey string
	}{
		Title:      "Contact Us",
		Path:       r.URL.Path,
		CaptchaKey: os.Getenv("HCAPTCHA_SITE_KEY"),
	}

	lp, hp := "templates/layout.gohtml", "templates/contact.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	defer tmpl.ExecuteTemplate(w, "layout", &data)

	if r.Method == http.MethodPost {
		values := struct{ Name, Email, Message, Token string }{
			r.FormValue("name"),
			r.FormValue("email"),
			r.FormValue("message"),
			r.FormValue("h-captcha-response"),
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

		to := mail.NewEmail("Enquiries", "enquiries@hollyshatchlings.co.uk")
		subject := "Website Enquiry from " + values.Name

		basePlain := "Name: %v\r\nEmail: %v\r\nMessage: %v\r\n"
		baseHTML := strings.TrimSpace(`
<div><b style="color:orange">Name:</b> %v</div>
<div><b style="color:orange">Email:</b> %v</div>
<div><b style="color:orange">Message:</b> %v</div>
		`)

		contentPlain := fmt.Sprintf(basePlain, values.Name, values.Email, values.Message)
		contentHTML := fmt.Sprintf(baseHTML, values.Name, values.Email, values.Message)

		err = mail.Send(to, subject, contentPlain, contentHTML)
		if err == nil {
			data.Success = true
		}
	}
}
