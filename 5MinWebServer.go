package main

import (
	"log"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
)

type Person struct {
	UserName string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.tmpl"))  // Parse template file.
	person := Person{UserName: "Astaxie"}
	t.Execute(w, person)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/", handler)

	err := http.ListenAndServeTLS(":9999", "cert.pem", "key.pem", router)
	if err != nil {
		log.Printf("Web Server cannot start: %v\n", err)
	}
}