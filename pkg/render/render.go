package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ShehabEl-DeenAlalkamy/residencia/pkg/config"
	"github.com/ShehabEl-DeenAlalkamy/residencia/pkg/models"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// perform some logic

	return td
}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// create a template cache
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from the template cache")
	}

	// create a buffer of bytes to try to execute the value we got from the map t
	// rather than doing it directly, we are going to execute the buffer directly and then write it out. The reason to do this is for finer grained error checking
	// because once I have declared this buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// this gives me a clear indication that the value I got from that map, there is something wrong with it. it parsed it, but I can't execute it
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
	// both the following lines are valid but the uncommented one is the most common one to create an empty map
	// myCache := make(map[string]template.Template)
	myCache := map[string]*template.Template{}

	// bear in mind when you are rendering a template that uses a layout, you typically must have as the first thing you try to parse the template you want to render and then the associated layouts
	// and partials after the fact

	// get all of the files manmed *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// ex: "./templates/home.page.tmpl" --> "home.page.tmpl"
		name := filepath.Base(page)

		// ts --> template set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look for any layouts that exist in that directory
		// 'matches' he meant 'layouts'
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// parses the templates defined in *.layout.tmpl and associates the result with ts
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}