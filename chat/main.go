package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth/providers/google"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/signature"
)

var (
	googleClientID     = os.Getenv("Google_Client_ID")
	googleClientSecret = os.Getenv("Google_Client_Secret")
)

var avatar Avatar = UseFileSystemAvatar

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if cookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(cookie.Value)
	}
	t.templ.Execute(w, data)
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
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))
	go r.run()
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalln("ListenAndServe error:", err)
	}
}
