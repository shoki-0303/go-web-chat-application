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
