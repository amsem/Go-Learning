package main

import ("net/http"
	"fmt"
	"html/template"
	"github.com/amsem/museum/data"
	"github.com/amsem/museum/api"
	)

var i int
func handleHome(w http.ResponseWriter, r *http.Request) {
                i++
                fmt.Printf("%d Access to the route /home\n",i)
                w.Write([]byte("Hello from my new home"))
        }

func handleTemplate(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/index.tmpl")
		if err !=nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		html.Execute(w, data.GetAll())
	}

func main() {
	server := http.NewServeMux()
	fmt.Println("Server running on route: /home port:5555")
	server.HandleFunc("/home", handleHome)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibition", api.Get)
	server.HandleFunc("/api/exhibition/post", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":5555", server)
	if err == nil {
		fmt.Println("Error while running the server")
	}
}

