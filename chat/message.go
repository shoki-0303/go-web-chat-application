package main

import (
	"time"
)

type message struct {
	Time      time.Time
	Message   string
	Name      string
	AvatarURL string
}
