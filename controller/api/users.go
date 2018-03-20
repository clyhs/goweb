package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/clyhs/goweb/controller/models"
)

func (a *Api) User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	log.Infof("......users.go.....")
	user, err := models.GetUserByName("cly")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
