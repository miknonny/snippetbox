package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html"}

	// parse the template.
	ts, err := template.ParseFiles(files...)
	// check for errors.
	if err != nil {

		// this error is for the app developer.
		app.serverError(w, err)
		return
	}

	// write the template to client and check for errors if any.
	if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)

		// simliar to the above code.
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a new snippet"))
}
