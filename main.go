package main

import (
	"html/template"
	"net/http"
	"time"
)

var (
	tpl = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/projects/", projects)
	http.ListenAndServe(":8080", nil)
}

func projects(w http.ResponseWriter, req *http.Request) {
	headfoot(w, req)
	tpl.ExecuteTemplate(w, "projects.gohtml", time.Now().Year())

}

func index(w http.ResponseWriter, req *http.Request) {
	headfoot(w, req)
	images := [5]string{"mountains", "cloudExpload", "MRC_1", "MRC_2", "MRC_3"}
	tpl.ExecuteTemplate(w, "index.gohtml", images)
}

func headfoot(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "header.gohtml", nil)
	tpl.ExecuteTemplate(w, "footer.gohtml", nil)

}
