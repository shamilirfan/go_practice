package handlers

import (
    "html/template"
    "net/http"
)

func ResetPasswordForm(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    tmpl := template.Must(template.ParseFiles("templates/reset_form.html"))
    tmpl.Execute(w, token)
}
