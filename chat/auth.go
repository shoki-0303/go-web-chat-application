package main

import (
	"fmt"
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
		//未認証
		fmt.Println("未認証")
	} else if err != nil {
		panic(err.Error())
	} else {
		//認証
		fmt.Println("認証")
	}
}
