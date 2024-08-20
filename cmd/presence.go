package main

import (
	"errors"
	"sync"
	"time"
)

func newPresence() *Presence {
	var lock sync.Mutex
	users := make(map[string]*PresenceDetails)
	return &Presence{lock: &lock, Users: users}
}

type PresenceDetails struct {
	UserID      string `json:"id"`
	Name        string `json:"name"`
	login       time.Time
	lastUpdate  time.Time
	LastDrawing string `json:"last_drawing"`
}

type Presence struct {
	lock  *sync.Mutex
	Users map[string]*PresenceDetails
}

func (p *Presence) AddUser(user *PresenceDetails) {
	p.lock.Lock()
	defer p.lock.Unlock()

	_, found := p.Users[user.UserID]
	if found {
		return
	}

	p.Users[user.UserID] = user
	user.login = time.Now()
	user.lastUpdate = time.Now()
}

func (p *Presence) IsUserPresent(userID string) bool {
	_, found := p.Users[userID]
	return found
}

func (p *Presence) RemoveUser(user *PresenceDetails) {}

func (p *Presence) UserAtDrawing(userID, drawingID string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	user, found := p.Users[userID]
	if !found {
		return errors.New("user not found")
	}

	user.lastUpdate = time.Now()
	user.LastDrawing = drawingID
	return nil
}
