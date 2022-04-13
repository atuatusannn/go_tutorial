package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"gopkg.in/go-ini/ini.v1/config"
)

func generateHTML(w http.ResponseWriter, date interface{}, filename ...string) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", date)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
