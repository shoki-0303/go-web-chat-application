package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
)

type authHandler struct {
	next http.Handler
}

func mustAuth(h http.Handler) http.Handler {
	return &authHandler{next: h}
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Println("login-provider error", err)
			return
		}
		url, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Println("login-geturl error", err)
			return
		}

		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		fmt.Println("callback時の処理を記述")
	default:
		w.WriteHeader(http.StatusNotFound)
		log.Printf("%sという誤ったアクションを検出しました", action)
	}
}
