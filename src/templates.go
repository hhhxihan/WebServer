package src

import(
	"html/template"
	
)

var Templates *template.Template

func LoadWeb(webRoute string){
	Templates=template.Must(template.ParseGlob(webRoute))
}