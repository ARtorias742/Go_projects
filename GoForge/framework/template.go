package framework

import (
    "html/template"
    "net/http"
    "path/filepath"
)

type TemplateRenderer struct {
    templates *template.Template
}

func NewTemplateRenderer(dir string) (*TemplateRenderer, error) {
    tmpl := template.New("").Funcs(template.FuncMap{
        "button": func(text string) template.HTML {
            return template.HTML(`<button class="btn">` + text + `</button>`)
        },
    })
    tmpl, err := tmpl.ParseGlob(filepath.Join(dir, "*.html"))
    if err != nil {
        return nil, err
    }
    tmpl, err = tmpl.ParseGlob(filepath.Join(dir, "components", "*.html"))
    if err != nil {
        return nil, err
    }
    return &TemplateRenderer{templates: tmpl}, nil
}

func (tr *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) error {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    return tr.templates.ExecuteTemplate(w, name, data)
}