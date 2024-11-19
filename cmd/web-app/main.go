package main

import (
	"github.com/guillechuma/url-shortener-app/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
