package main

import (
	"k8s-test/handlers"
	"k8s-test/middlewares"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	const port = ":8000"

	var (
		username         = os.Getenv("USERNAME")
		password         = os.Getenv("PASSWORD")
		itemsQuantity, _ = strconv.Atoi(os.Getenv("ITEMS_QUANTITY"))
	)

	http.Handle("/", middlewares.TrackRequest(handlers.Home(username, password)))
	http.HandleFunc("/items", handlers.Items(itemsQuantity))

	log.Printf("Server started at %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
