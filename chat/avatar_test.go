package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)

	avatarURL, err := authAvatar.GetAvatar(client)
	if err != ErrNoAvatar {
		t.Error("GetAvatar should return ErrNoAvatar")
	}
	testURL := "http://avatar-test-url"
	client.userData = map[string]interface{}{
		"avatar_url": testURL,
	}
	avatarURL, err = authAvatar.GetAvatar(client)
	if err != nil {
		t.Error("GetAvatar should not return error, when avatar-url exists")
	} else if avatarURL != testURL {
		t.Error("GetAvatar should return right url")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)

	client.userData = map[string]interface{}{
		"userid": "abc",
	}
	url, err := gravatarAvatar.GetAvatar(client)
	if err != nil {
		t.Error("GetAvatar should not return error")
	} else if url != "https://www.gravatar.com/avatar/abc" {
		t.Error("GetAvatar should return right url")
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpeg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "abc",
	}

	url, err := fileSystemAvatar.GetAvatar(client)
	if err != nil {
		log.Println(err)
		t.Error("GetAvatar should not return error")
	} else if url != "/avatars/abc.jpeg" {
		t.Error("GetAvatar should return right url")
	}

}
