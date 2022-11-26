package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string
	Description string
	CreateOn    time.Time
}

var noteStore = make(map[string]Note)
var id int = 0

var templates (map[string]*template.Template)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)

	}
	templates["index"] = template.Must(template.ParseFiles("template/index.html", "template/base.html"))
	templates["add"] = template.Must(template.ParseFiles("template/add.html", "template/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("template/edit.html", "template/base.html"))
}

func renderTemplate(rw http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(rw, "The template does not exist.", http.StatusInternalServerError)
		return
	}

	err := tmpl.ExecuteTemplate(rw, template, viewModel)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func getNotes(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "Index", "base", noteStore)
}

func addNote(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "add", "base", nil)
}

func saveNote(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	note := Note{title, description, time.Now()}
	id++

	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(rw, r, "/", 302)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	// r.HandleFunc("/notes/edit/{id}", editNote)
	// r.HandleFunc("/notes/update/{id}", updateNote)
	// r.HandleFunc("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	fmt.Println("Listning at port 8080")
	server.ListenAndServe()

}
