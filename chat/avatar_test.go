package main

import (
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	chatuser := new(chatuser)

	avatarURL, err := authAvatar.GetAvatar(chatuser)
	if err != ErrNoAvatar {
		t.Error("GetAvatar should return ErrNoAvatar")
	}

	TestURL := "http://avatar-url"
	chatuser.avatarurl = TestURL

	avatarURL, err = authAvatar.GetAvatar(chatuser)
	if err != nil {
		t.Error("GetAvatar should not return error")
	} else if avatarURL != TestURL {
		t.Error("GetAvatar should return right url")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	chatuser := new(chatuser)
	chatuser.userid = "abc"

	avatarURL, err := gravatarAvatar.GetAvatar(chatuser)
	if err != nil {
		t.Error("GetAvatar should not return error")
	} else if avatarURL != "https://www.gravatar.com/avatar/abc" {
		t.Error("GetAvatar should return right url")
	}
}

func TestFileSystemAvatar(t *testing.T) {
	t.Skip("skip now")
}
