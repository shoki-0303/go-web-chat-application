package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatar : ErrNoAvatar returns error, when avatar instance can't return url.
var ErrNoAvatar = errors.New("avatar-url doesn't exist")

// Avatar : avatar instance has a GetAvatar method
type Avatar interface {
	GetAvatar(u *chatuser) (string, error)
}

// AuthAvatar : get avatar via auth
type AuthAvatar struct{}

// UseAuthAvatar : instance of AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatar : get avatar-url
func (AuthAvatar) GetAvatar(u *chatuser) (string, error) {
	if u.getAvatarURL() != "" {
		return u.getAvatarURL(), nil
	}
	return "", ErrNoAvatar
}

// GravatarAvatar : get avatar via gravatar
type GravatarAvatar struct{}

// UseGravatarAvatar : instance of GravatarAvatar
var UseGravatarAvatar GravatarAvatar

// GetAvatar : get avatar-url
func (GravatarAvatar) GetAvatar(u *chatuser) (string, error) {
	if u.getUserID() != "" {
		return "https://www.gravatar.com/avatar/" + u.getUserID(), nil
	}
	return "", ErrNoAvatar
}

// FileSystemAvatar : get avatar via local-file system
type FileSystemAvatar struct{}

// UseFileSystemAvatar : instance of FileSystemAvatar
var UseFileSystemAvatar FileSystemAvatar

// GetAvatar : get avatar-url
func (FileSystemAvatar) GetAvatar(u *chatuser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if match, _ := filepath.Match(u.getUserID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatar
}
