package main

import (
	"log"
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func mustAuth(h http.Handler) http.Handler {
	return &authHandler{next: h}
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		log.Println("未認証")
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		log.Println("認証通過")
		h.next.ServeHTTP(w, r)
	}
}
