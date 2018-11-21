package main

import (
	"errors"
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
