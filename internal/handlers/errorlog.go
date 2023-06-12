package handlers

import (
	"log"
	"net/http"
)

func errorlog(w http.ResponseWriter, errText string, errStatus int) {
	http.Error(w, errText, errStatus)
	log.Println("Error: ", errText)
}
