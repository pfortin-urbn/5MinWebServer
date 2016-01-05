package main

import (
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
	router.HandleFunc("/", handler)

	http.ListenAndServe(":9999", router)
}