package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ethanjmarchand/rewtwo/pkg/config"
	"github.com/ethanjmarchand/rewtwo/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData makes all the data available to all templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	// buf simply creates a space to store bytes
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	// instead of writing to "w", execute is writing to my storage space "buf" to first check for errors
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all of the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all strings stored in Pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			// Parseglob is taking the above "ts"... an additional all of the layout files to the parse.
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}

// var tc = map[string]*template.Template{}

// func RenderTemplateSimpleCache(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
// 	// check to see if we already have the template saved in tc (template cache)
// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create tc (template cache)
// 		log.Println("Creating template, and adding to tc")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have the template in tc (template cache)
// 		log.Println("using tc")
// 	}
// 	// by now, we have tmpl
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	// add template to cache
// 	tc[t] = tmpl
// 	return nil
// }
