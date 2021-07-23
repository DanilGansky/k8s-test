package handlers

import (
	"encoding/json"
	"k8s-test/storage"
	"k8s-test/utils"
	"log"
	"net/http"
)

func Items(itemsQuantity int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(storage.GetItems(itemsQuantity)); err != nil {
			log.Fatalf("error in %s: %s", utils.GetOutboundIP(), err.Error())
		}
	}
}
