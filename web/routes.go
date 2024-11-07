package web

import (
    "html/template"
    "net/http"
    "log"
)

var tmpl *template.Template

type Person struct {
    FirstName string
    LastName string
}

func init() {
    var err error
    tmpl, err = template.ParseGlob("web/templates/*.html")
    if err!=nil {
        log.Fatalf("Error serving templates: %v", err)
    }
}

func Routes() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeRouteHandler)
    mux.HandleFunc("/about", aboutRouteHandler)
    return mux
}

func homeRouteHandler(w http.ResponseWriter, r *http.Request){
    var data = Person{"Nishant", "Chhillar"}
    tmpl.ExecuteTemplate(w, "index.html", data)
}

func aboutRouteHandler(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "about.html", nil)
}
