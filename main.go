package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./public/*.html",
	}

	rnd = renderer.New(opts)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s at %s", r.Host, time.Now().Format("Mon Jan 2 15:04:05 2006"))

	err := rnd.Template(w, http.StatusOK, []string{"./public/index.tpls"}, map[string]interface{}{"ip": r.Host})
	if err != nil {
		log.Fatalf("error in %s: %s", r.RemoteAddr, err.Error())
	}
}

func main() {
	const port = ":8000"

	http.HandleFunc("/", homeHandler)
	log.Printf("Server started at %s...", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
