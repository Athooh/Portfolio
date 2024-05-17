package handlers

import (
	"log"
	"net/http"
	"portfolio/common"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	err := common.Tpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
