package main

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"rentalgen/config"
	"rentalgen/logger"
	"rentalgen/parser"
)

func main() {
	flags := config.GetFlags()

	tmplData, err := parser.Parse(flags)
	if err != nil {
		logger.Logln(logger.ColorRed, err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(
			filepath.Join("templates", "root.html"),
			filepath.Join("templates", "voucher.html"),
			filepath.Join("templates", "bordereau.html"),
		)

		tmpl.ExecuteTemplate(w, "root", tmplData)
	})

	http.ListenAndServe(":8080", mux)
}
