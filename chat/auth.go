package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
)

// ChatUser : has getAvatarURL method and getUserID method
type ChatUser interface {
	getAvatarURL() string
	getUserID() string
}

type chatuser struct {
	avatarurl string
	userid    string
}

func (u chatuser) getAvatarURL() string {
	return u.avatarurl
}
func (u chatuser) getUserID() string {
	return u.userid
}

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
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Println("callback-provider error", err)
			return
		}
		data := objx.MustFromURLQuery(r.URL.RawQuery)
		creds, err := provider.CompleteAuth(data)
		if err != nil {
			log.Println("callback-creds error", err)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Println("callback-user error", err)
			return
		}

		chatUser := &chatuser{
			avatarurl: user.AvatarURL(),
		}
		hasher := md5.New()
		io.WriteString(hasher, strings.ToLower(user.Email()))
		userid := fmt.Sprintf("%x", hasher.Sum(nil))
		chatUser.userid = userid

		avatarURL, err := avatars.GetAvatar(chatUser)
		if err != nil {
			log.Println(err)
			return
		}

		authCookieValue := objx.New(map[string]interface{}{
			"name":       user.Name(),
			"avatar_url": avatarURL,
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Path:  "/",
			Value: authCookieValue,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		log.Printf("%sという誤ったアクションを検出しました", action)
	}
}
