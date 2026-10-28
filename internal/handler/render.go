package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	templates map[string]*template.Template
}

func NewRenderer() *Renderer {
	tmplMap := make(map[string]*template.Template)
	baseLayout := "web/templates/layouts/base.html"

	pages := []string{
		"public/index.html",
		"public/blog.html",
		"public/post.html",
		"admin/list.html",
		"admin/form.html",
	}

	for _, page := range pages {
		fullPath := filepath.Join("web/templates", page)
		tmplMap[page] = template.Must(template.ParseFiles(baseLayout, fullPath))
	}

	return &Renderer{templates: tmplMap}
}

func (rn *Renderer) Render(w http.ResponseWriter, pageName string, data any) {
	tmpl, exists := rn.templates[pageName]
	if !exists {
		http.Error(w, "قالب پیدا نشد: "+pageName, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "base", data)
}
