package handlers

import (
	"html/template"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/gorm/clause"
)

// Animal returns a page for a single animal.
func Animal(w http.ResponseWriter, r *http.Request) {
	pageData := struct {
		Title, Path string
		Animal      *model.Animal
	}{
		Path: r.URL.Path,
	}

	var animal model.Animal
	database.DB.Model(&model.Animal{}).
		Preload(clause.Associations).
		Where("id = ?", mux.Vars(r)["id"]).
		First(&animal)

	if reflect.ValueOf(animal).IsZero() {
		NotFound(w, r)
		return
	}

	animal.LoadImages()
	pageData.Animal = &animal

	var isPersonal bool
	for _, status := range []string{"Future Breeder", "Breeder", "Non Breeder"} {
		if status == animal.Status {
			isPersonal = true
			break
		}
	}

	if len(animal.Name) > 0 && isPersonal {
		pageData.Title = animal.Name
	} else if len(animal.Reference) > 0 {
		pageData.Title = animal.Reference
	} else {
		pageData.Title = animal.Species.Name
	}

	lp, hp := "templates/layout.gohtml", "templates/animal.gohtml"
	tmpl := template.New("Animal").Funcs(template.FuncMap{
		"formatDate": func(t *time.Time) string {
			return t.Format("02/01/2006")
		},
	})
	tmpl = template.Must(tmpl.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pageData)
}
