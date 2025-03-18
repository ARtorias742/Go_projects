package framework

import "net/http"

func ParseForm(r *http.Request) map[string]string {
	r.ParseForm()
	formData := make(map[string]string)
	for key, values := range r.Form {
		if len(values) > 0 {
			formData[key] = values[0]
		}
	}
	return formData
}
