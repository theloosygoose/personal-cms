package utils

import (
    "github.com/a-h/templ"
    "net/http"
)

func TemplRender(w http.ResponseWriter, r *http.Request , component templ.Component) error {
    return component.Render(r.Context(), w)
}
