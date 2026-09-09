package main

import (
	"bufio"
	"crypto/tls"
	"time"
)

const (
	SpammerGoroutines = 25
	SpammerDuration   = time.Millisecond * 500
)

var (
	fullUsernames          []string
	usernames              []string
	usernameIDs            []string
	usernameBodies         []string
	usernameMatches        [][]byte
	activeSessions         []string
	activeSession          []string
	spammerUsernameIndex   int = -1
	spammerDummyRequest    []byte
	usernameChangeRequests [][]byte
	usernameChangeRequest  []byte
	globalConnection       *tls.Conn
	globalBuffer           *bufio.Reader
)
