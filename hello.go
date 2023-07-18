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
	fmt.Println("Hello, World!")

	// Handle #1 - index request and return "Hello World"
	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "Hello World\n")
	// 	io.WriteString(w, "httpMethod : "+r.Method)
	// }

	// Handle #2 - index request and return index.html template, with film data
	h2 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Fancis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	// Handler #3 - return template block with newly added film/director, as an HTMX response
	h3 := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request received")
		// time sleep just to display spinner effect on submit button
		time.Sleep(1 * time.Second)

		// refer via html name
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		fmt.Println(title)
		fmt.Println(director)

		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h2)
	http.HandleFunc("/add-film/", h3)

	// Host as server on 8080 port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
