package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/raziel2244/geckosite/mail"
)

// Contact returns the contact us page.
func Contact(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Success bool
		Message string
	}{}

	if r.Method == http.MethodPost {
		basePlain := "Name: %v\r\nEmail: %v\r\nSubject: %v\r\nMessage: %v\r\n"

		baseHTML := strings.TrimSpace(`
<div><b style="color:orange">Name:</b> %v</div>
<div><b style="color:orange">Email:</b> %v</div>
<div><b style="color:orange">Message:</b> %v</div>
		`)

		values := struct{ Name, Email, Message string }{
			r.FormValue("name"),
			r.FormValue("email"),
			r.FormValue("message"),
		}

		to := mail.NewEmail("Enquiries", "enquiries@hollyshatchlings.co.uk")
		subject := "Website Enquiry from " + values.Name

		contentPlain := fmt.Sprintf(basePlain, values.Name, values.Email, values.Message)
		contentHTML := fmt.Sprintf(baseHTML, values.Name, values.Email, values.Message)

		err := mail.Send(to, subject, contentPlain, contentHTML)
		if err == nil {
			data.Success = true
		}
	}

	lp, hp := "templates/layout.gohtml", "templates/contact.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
