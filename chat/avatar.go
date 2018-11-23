package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// ErrNoAvatar : ErrNoAvatar returns error, when avatar instance can't return url.
var ErrNoAvatar = errors.New("avatar-url doesn't exist")

// Avatar : avatar instance has a GetAvatar method
type Avatar interface {
	GetAvatar(c *client) (string, error)
}

// AuthAvatar : get avatar via auth
type AuthAvatar struct{}

// UseAuthAvatar : instance of AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatar : get avatar-url
func (AuthAvatar) GetAvatar(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatar
}

// GravatarAvatar : get avatar via gravatar
type GravatarAvatar struct{}

// UseGravatarAvatar : instance of GravatarAvatar
var UseGravatarAvatar GravatarAvatar

// GetAvatar : get avatar-url
func (GravatarAvatar) GetAvatar(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			emailStrLow := strings.ToLower(emailStr)
			hasher := md5.New()
			io.WriteString(hasher, emailStrLow)
			url := fmt.Sprintf("https://www.gravatar.com/avatar/%x", hasher.Sum(nil))
			return url, nil
		}
	}
	return "", ErrNoAvatar
}

// FileSystemAvatar : get avatar via local-file system
type FileSystemAvatar struct{}

// UseFileSystemAvatar : instance of FileSystemAvatar
var UseFileSystemAvatar FileSystemAvatar

// GetAvatar : get avatar-url
func (FileSystemAvatar) GetAvatar(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if match, _ := filepath.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatar
}
