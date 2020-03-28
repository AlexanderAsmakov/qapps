package handlers

import (
	"fmt"
	"github.com/AlexanderAsmakov/qapps/internal/models"
	"net/http"
)

var GetAllApps = func(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	apps := models.GetAllApps()

	for _, app := range apps {
		fmt.Fprintf(w, "%d, %s, %s, %s, %s\n", app.ID, app.Name, app.Title, app.BundleIdentifier, app.OS)
	}
}
