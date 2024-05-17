package handlers

import (
	"log"
	"net/http"
	"portfolio/common"
)

func About(w http.ResponseWriter, r *http.Request) {
	err := common.Tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
