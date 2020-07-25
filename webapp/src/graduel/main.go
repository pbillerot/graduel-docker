package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Todo ...
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData ...
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var templates = template.Must(template.ParseFiles(
	"templates/home.html",
	"templates/base.html"))

func main() {

	http.HandleFunc("/", homeHandler)

	log.Println("Listening on http://localhost:3003...")
	err := http.ListenAndServe(":3003", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err := templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func listDir() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Println(file)
	}
}
