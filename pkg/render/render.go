package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Coding-Bruh/bookings/pkg/config"
	"github.com/Coding-Bruh/bookings/pkg/models"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, htmlTmpl string, td *models.TemplateData) {
	// new implementation get the template cache from app.config
	var tc map[string]*template.Template
	if app.UseCache {
		// get the templacache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//// create a template cache
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// get requested template from cache
	t, ok := tc[htmlTmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	er := t.Execute(buf, td)
	if er != nil {
		log.Println(er)
	}

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*tempplate.Template) // but use other way shown below
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.path.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
