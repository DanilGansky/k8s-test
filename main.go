package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/thedevsaddam/renderer"
)

var (
	rnd     *renderer.Render
	localIP string
)

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./public/*.html",
	}

	rnd = renderer.New(opts)
	localIP = getOutboundIP().String()
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s at %s", localIP, time.Now().Format("Mon Jan 2 15:04:05 2006"))

	err := rnd.Template(w, http.StatusOK, []string{"./public/index.tpls"}, map[string]interface{}{"ip": localIP})
	if err != nil {
		log.Fatalf("error in %s: %s", localIP, err.Error())
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
