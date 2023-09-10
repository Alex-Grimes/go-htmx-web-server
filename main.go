package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	fmt.Println("server is running ...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
				{Title: "Bullitt", Director: "Peter Yates"},
				{Title: "The French Connection", Director: "William Friedkin"},
			},
		}

		tmpl.Execute(w, films)
	}

	addFilm := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-elemet", Film{Title: title, Director: director})

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", addFilm)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
