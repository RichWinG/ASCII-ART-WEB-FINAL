package main

import (
	"ASCII-ART-WEB/pkg/internals/app"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func errHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case 404:
		ln, err := template.ParseFiles("./templates/notfound.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		w.WriteHeader(http.StatusNotFound)
		ln.Execute(w, nil)
	case 500:
		ln, err := template.ParseFiles("./templates/internal.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		ln.Execute(w, nil)
	case 400:
		ln, err := template.ParseFiles("./templates/badrequest.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		ln.Execute(w, nil)
	}
}

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	ln, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError)
		return
	}
	ln.Execute(w, nil)
}

// Обработчик для отображения содержимого заметки.
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		errHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		errHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banners := r.FormValue("banners")
	text = strings.ReplaceAll(text, "\r", "")
	result, er := app.Run(text, banners)
	if er != 0 {
		if er == 400 {
			errHandler(w, r, 400)
			return
		}
		if er == 500 {
			errHandler(w, r, 500)
			return
		}
	}
	ln, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError)
	}
	ln.Execute(w, result)
}

// Обработчик для создания новой заметки.
func createAscii(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", asciiHandler)
	file := http.FileServer(http.Dir("./templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates", file))
	log.Println("Starting the web server on localhost:5500")
	err := http.ListenAndServe(":5500", mux)
	log.Fatal(err)
}
