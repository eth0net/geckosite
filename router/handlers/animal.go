package handlers

import (
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/gorm/clause"
)

// Animal returns a page for a single animal.
func Animal(w http.ResponseWriter, r *http.Request) {
	var data model.Animal
	database.DB.Model(&model.Animal{}).
		Preload(clause.Associations).
		Where("id = ?", mux.Vars(r)["id"]).
		First(&data)

	if reflect.ValueOf(data).IsZero() {
		NotFound(w, r)
		return
	}

	lp, hp := "templates/layout.gohtml", "templates/animal.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
