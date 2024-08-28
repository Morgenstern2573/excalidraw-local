package main

import (
	"sync"
	"time"
)

type LockDetails struct {
	UserID    string
	DrawingID string
	LockedAt  time.Time
}

type DrawingLock struct {
	lock           *sync.Mutex
	LockedDrawings map[string]*LockDetails
}

func NewLock() *DrawingLock {
	var lock sync.Mutex

	dMap := make(map[string]*LockDetails, 0)

	return &DrawingLock{lock: &lock, LockedDrawings: dMap}
}

func (l *DrawingLock) LockDrawing(userID, drawingID string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	_, found := l.LockedDrawings[userID]

	if found {
		return
	}

	l.LockedDrawings[userID] = &LockDetails{
		UserID:    userID,
		DrawingID: drawingID,
		LockedAt:  time.Now(),
	}

}

func (l *DrawingLock) RemoveUser(userID string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	_, found := l.LockedDrawings[userID]

	if !found {
		return
	}

	delete(l.LockedDrawings, userID)
}

func (l *DrawingLock) IsDrawingLocked(drawingID string) (string, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for userID, details := range l.LockedDrawings {
		if details.DrawingID == drawingID {
			return userID, true
		}
	}

	return "", false
}
