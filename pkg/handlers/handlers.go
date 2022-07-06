package handlers

import (
	"net/http"

	"github.com/guimassoqueto/go_web_app/pkg/render"
)

// home controller
func Home(res http.ResponseWriter, req *http.Request) {
	// rendering home template
	render.RenderTemplate(res, "home.page.html")
}

// about controller
func About(res http.ResponseWriter, req *http.Request) {
	// rendering about template
	render.RenderTemplate(res, "about.page.html")
}
