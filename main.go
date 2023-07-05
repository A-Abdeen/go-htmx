package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// time.Sleep(1 * time.Second) <-- To simulate a loading buffer on a POST method handler

/*
Add context to template:
1. Create new type (Film) to group data & form records
2. Make a list of films
3. Pass data to template to execute
4. Access data within the template's body tag to render for client
*/
type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	fmt.Println("localhost:8080 is now active")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
