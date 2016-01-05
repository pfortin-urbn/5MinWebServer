package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Hello World!!!"))
}

func main() {
router := mux.NewRouter().StrictSlash(true)

router.HandleFunc("/", handler)
http.ListenAndServe(":9999", router)
}

