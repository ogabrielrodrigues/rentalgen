package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"

	"rentalgen/config"
	"rentalgen/logger"
	"rentalgen/parser"
	"rentalgen/templates"
)

const addr = "127.0.0.1:8080"

func main() {
	flags := config.GetFlags()

	tmplData, err := parser.Parse(flags.FilePath)
	if err != nil {
		fatal(err)
	}

	tmpl, err := template.ParseFS(templates.FS, "*.html")
	if err != nil {
		fatal(err)
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fatal(fmt.Errorf("✗ erro ao iniciar o servidor: %w", err))
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		var page bytes.Buffer

		err := tmpl.ExecuteTemplate(&page, "root", tmplData)
		if err != nil {
			logger.Errorln(err.Error())
			http.Error(w, "erro ao gerar recibos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page.WriteTo(w)
	})

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, templates.FS, "styles.css")
	})

	logger.Log(logger.ColorGreen, "✓ arquivo de alugueres selecionado: ")
	logger.Logln(logger.ColorDefault, flags.FilePath)
	logger.Logln(logger.ColorDefault, "✓ para realizar a impressão dos recibos, acesse http://localhost:8080\n")
	logger.Logln(logger.ColorDefault, "Para encerrar, pressione Ctrl + C")

	err = http.Serve(listener, mux)
	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	logger.Errorln(err.Error())
	os.Exit(1)
}
