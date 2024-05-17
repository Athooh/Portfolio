package common 

import (
	"html/template"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("./template/*.html"))
}