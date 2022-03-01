package render

import (
	"bytes"
	"fmt"
	"github.com/williamneokh/booking/pkg/config"
	"github.com/williamneokh/booking/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var Functions = template.FuncMap{}

var app *config.Appconfig

func NewTemplate(a *config.Appconfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc = map[string]*template.Template{}
	if app.UserCache {
		tc = app.TemplateCache
	} else {
		tc = TemplateSet()
	}
	t := tc[tmpl]

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func TemplateSet() map[string]*template.Template {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Println("Error looking up for page template", err)
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(Functions).ParseFiles(page)

		if err != nil {
			log.Fatal(err)
		}
		ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

		if err != nil {
			log.Fatal(err)
		}
		myCache[name] = ts

	}
	return myCache

}
