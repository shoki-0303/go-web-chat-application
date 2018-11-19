package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	r := newRoom()
	addr := flag.String("addr", ":8080", "address of application")
	flag.Parse()
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalln("ListenAndServe error:", err)
	}
}
