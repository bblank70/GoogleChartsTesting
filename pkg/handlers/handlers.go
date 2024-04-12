package handlers

import (
	"net/http"

	"github.com/bblank70/GoogleChartsTesting/pkg/render"
)

// index is the home page handler
func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "base.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}
