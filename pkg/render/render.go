package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(res, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
