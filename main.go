package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Album struct {
	Band string
	Name string
}

func main() {
	fmt.Println("hello")
	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))

		content := map[string][]Album{
			"Albums": {
				{Band: "Nirvana", Name: "In utero"},
				{Band: "Love", Name: "Forever changes"},
			},
		}

		tmpl.Execute(w, content)
	}
	js := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
