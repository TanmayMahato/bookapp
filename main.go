package main

import (
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	BookName   string
	AuthorName string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		books := map[string][]Book{
			"Books": {
				{BookName: "Anxious People", AuthorName: "Fredrick Backman"},
				{BookName: "Anxious People", AuthorName: "Fredrick Backman"},
				{BookName: "Anxious People", AuthorName: "Fredrick Backman"},
			}}

		tmpl.Execute(w, books)

	}
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		b := r.PostFormValue("bookName")
		a := r.PostFormValue("authorName")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "book-element", Book{BookName: b, AuthorName: a})

	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/recbook/", handler2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
