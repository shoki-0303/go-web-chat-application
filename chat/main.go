package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth/providers/google"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/signature"
)

var (
	googleClientID     = os.Getenv("Google_Client_ID")
	googleClientSecret = os.Getenv("Google_Client_Secret")
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
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:8080/auth/callback/google"),
	)
	http.Handle("/chat", mustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	go r.run()
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalln("ListenAndServe error:", err)
	}
}
