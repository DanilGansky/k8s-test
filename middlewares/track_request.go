package middlewares

import (
	"k8s-test/utils"
	"log"
	"net/http"
	"time"
)

func TrackRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request from %s at %s", utils.GetOutboundIP(), time.Now().Format(time.RFC822))
		next.ServeHTTP(w, r)
	}
}
