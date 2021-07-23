package handlers

import (
	"k8s-test/utils"
	"log"
	"net/http"
)

func Home(username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		localIP := utils.GetOutboundIP()
		renderer := utils.GetRenderer()
		templatePath := []string{"./public/index.tpls"}
		context := map[string]interface{}{
			"ip":       localIP,
			"username": username,
			"password": password,
		}

		if err := renderer.Template(w, http.StatusOK, templatePath, context); err != nil {
			log.Fatalf("error in %s: %s", localIP, err.Error())
		}
	}
}
