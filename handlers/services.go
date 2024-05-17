package handlers

import (
	"net/http"
	"portfolio/common"
	"log"
)


func Services(w http.ResponseWriter, r *http.Request) {
	err := common.Tpl.ExecuteTemplate(w, "services.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}