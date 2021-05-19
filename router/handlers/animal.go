package handlers

import (
	"html/template"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"github.com/raziel2244/geckosite/templates"
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

	pageData.Animal = &animal

	var isPersonal bool
	for _, status := range []string{"Future Breeder", "Breeder", "Non Breeder"} {
		if status == animal.Status {
			isPersonal = true
			break
		}
	}

	if isPersonal && animal.Name.Valid && len(animal.Name.String) > 0 {
		pageData.Title = animal.Name.String
	} else if animal.Reference.Valid && len(animal.Reference.String) > 0 {
		pageData.Title = animal.Reference.String
	} else {
		pageData.Title = animal.Species.Name
	}

	tmpl := template.New("Animal").Funcs(template.FuncMap{
		"formatDate": func(t *time.Time) string {
			return t.Format("02/01/2006")
		},
	})
	tmpl = templates.ParseInto(tmpl, "animal")
	tmpl.ExecuteTemplate(w, "layout", pageData)
}
