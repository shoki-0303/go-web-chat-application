package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func uploaderHandler(w http.ResponseWriter, r *http.Request) {
	userid := r.FormValue("userid")
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		log.Println("FormFile", "-", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicln("ReadAll", "-", err)
		return
	}
	filename := userid + filepath.Ext(header.Filename)
	avatarFile := filepath.Join("avatars/", filename)
	err = ioutil.WriteFile(avatarFile, data, 0777)
	if err != nil {
		log.Println("WriteFile", "-", err)
		return
	}
	io.WriteString(w, "画像のホスティングに成功しました")
}
