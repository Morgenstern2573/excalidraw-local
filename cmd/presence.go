package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
)

func newPresence() *Presence {
	var lock sync.Mutex
	users := make(map[string]*PresenceDetails)
	return &Presence{lock: &lock, Users: users}
}

type PresencePosition struct {
	X string
	Y string
}

type PresenceDetails struct {
	UserID      string `json:"id"`
	Name        string `json:"name"`
	login       time.Time
	lastUpdate  time.Time
	LastDrawing string `json:"last_drawing"`
	Position    PresencePosition
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

func (p *Presence) RemoveUser(userID string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.Users, userID)
}

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

func (p *Presence) UpdateUserPosition(userID, drawingID string, position PresencePosition) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	details, found := p.Users[userID]
	if !found {
		return errors.New("user not found")
	}

	details.Position = position
	details.LastDrawing = drawingID
	details.lastUpdate = time.Now()

	return nil
}

func (p *Presence) GetUsersAtDrawing(drawingID string) ([]*PresenceDetails, error) {
	retv := make([]*PresenceDetails, 0)
	users := p.Users

	for _, details := range users {
		if details.LastDrawing == drawingID {
			retv = append(retv, details)
		}
	}
	return retv, nil
}

// TODO: GET rid of shitty name
func makePresenceMap(drawingList []services.Drawing, users map[string]*PresenceDetails, currentUser string) (map[string][]ui.PresentUser, error) {
	presenceMap := make(map[string][]ui.PresentUser)

	for _, details := range users {
		for _, drawing := range drawingList {
			if details.LastDrawing == drawing.ID {
				if details.UserID == currentUser {
					continue
				}

				user, err := services.Users().GetUserByID(details.UserID)
				if err != nil {
					return nil, err
				}

				initials := string(user.FirstName[0]) + string(user.LastName[0])
				name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
				userDetails := ui.PresentUser{Initials: initials, Name: name}

				_, found := presenceMap[details.LastDrawing]
				if found {
					presenceMap[details.LastDrawing] = append(presenceMap[details.LastDrawing], userDetails)
				} else {
					presenceMap[details.LastDrawing] = []ui.PresentUser{userDetails}
				}

				break
			}
		}
	}

	return presenceMap, nil
}
