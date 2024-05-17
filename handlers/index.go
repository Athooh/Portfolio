package handlers

import (
	"log"
	"net/http"
	"portfolio/common"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := common.Tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
