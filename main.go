package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Todo struct {
	Title       string
	Description string
}

func main() {

	fmt.Println("server is running ...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		todos := map[string][]Todo{
			"Todos": {
				{Title: "Walk The Dog", Description: "Take the dog for a walk"},
				{Title: "Walk The Dog", Description: "Take the dog for a walk"},
				{Title: "Walk The Dog", Description: "Take the dog for a walk"},
				{Title: "Walk The Dog", Description: "Take the dog for a walk"},
			},
		}

		tmpl.Execute(w, todos)
	}

	addTodo := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-elemet", Todo{Title: title, Description: description})

	}

	deleteTodo := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-elemet", Todo{Title: "", Description: ""})

	}

	updateTodo := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-elemet", Todo{Title: title, Description: description})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", addTodo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
