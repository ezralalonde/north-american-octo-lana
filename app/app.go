package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	templates = make(map[string]*template.Template)
)

func init() {
	templateInit()
	http.HandleFunc("/", rootHandler)
}

func templateInit() {
	load("demo")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t := get("demo")
		t.ExecuteTemplate(w, "root", nil)
	} else {
		w.Write([]byte(fmt.Sprintf("I'd rather be at %s", r.URL.Path)))
	}
}

func get(name string) *template.Template {
	if template, ok := templates[name]; ok {
		return template
	}
	load(name)
	return templates[name]
}

func load(name string) {
	fields := strings.Fields(name)
	shortname := fields[len(fields)-1]
	template := template.New(shortname)
	for _, name := range fields {
		if !strings.HasSuffix(name, ".html") {
			name += ".html"
		}
		if _, err := template.ParseFiles(name); err != nil {
			panic(err)
		}
	}
	templates[shortname] = template
}
