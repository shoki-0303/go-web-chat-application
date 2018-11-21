package main

import (
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
		"email": "MyExample@Gmail.com",
	}
	url, err := gravatarAvatar.GetAvatar(client)
	if err != nil {
		t.Error("GetAvatar should not return error")
	} else if url != "https://www.gravatar.com/avatar/62604ed7843f0e296d40bc82cb7342cc" {
		t.Error("GetAvatar should return right url")
	}
}
