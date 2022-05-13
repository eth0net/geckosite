package web

import (
	"html/template"
	"reflect"
	"time"

	"github.com/eth0net/geckosite/database/model"
	"github.com/eth0net/geckosite/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// Animal returns a page for a single animal.
func (s service) Animal(c *gin.Context) {
	pageData := struct {
		Title, Path string
		Animal      *model.Animal
	}{
		Path: c.Request.URL.Path,
	}

	var animal model.Animal
	s.db.Model(&model.Animal{}).
		Preload(clause.Associations).
		Where("id = ?", c.Param("id")).
		First(&animal)

	if reflect.ValueOf(animal).IsZero() {
		s.NotFound(c)
		return
	}

	pageData.Animal = &animal

	var isPersonal bool
	for _, status := range []string{"Future Breeder", "Breeder"} {
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
	tmpl.ExecuteTemplate(c.Writer, "layout", pageData)
}
