package handlers

import (
	"context"
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"github.com/raziel2244/geckosite/s3"
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

	ch := s3.Client.ListObjects(
		context.Background(),
		data.Species.Order,
		minio.ListObjectsOptions{
			Prefix:    data.Species.Type + "/" + data.ID.String(),
			Recursive: true,
		},
	)
	for object := range ch {
		path := "/s3/" + data.Species.Order + "/" + object.Key
		data.Images = append(data.Images, path)
	}

	lp, hp := "templates/layout.gohtml", "templates/animal.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
