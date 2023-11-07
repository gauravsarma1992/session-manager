package sessionmgmt

import (
	"errors"
	"sync"
)

type (
	SessionStore struct {
		store  map[TokenT]*Session
		config *SessionsStoreConfig

		storeLock *sync.RWMutex
	}

	SessionsStoreConfig struct {
		ActiveSessionsAllowed uint32
	}
)

var (
	SessionNotPresentError    = errors.New("No Session found")
	SessionAlreadyExistsError = errors.New("Session already exists")
)

func DefaultSessionStoreConfig() (sc *SessionsStoreConfig) {
	sc = &SessionsStoreConfig{
		ActiveSessionsAllowed: 1200,
	}
	return
}

func NewSessionStore(scStoreCfg *SessionsStoreConfig) (sStore *SessionStore, err error) {
	sStore = &SessionStore{
		store:     make(map[TokenT]*Session),
		config:    scStoreCfg,
		storeLock: &sync.RWMutex{},
	}
	if sStore.config == nil {
		sStore.config = DefaultSessionStoreConfig()
	}
	return
}

func (sStore *SessionStore) AddSession(token TokenT, obj *SessionObj) (session *Session, err error) {

	sStore.storeLock.Lock()
	defer sStore.storeLock.Unlock()

	// Check if the session with token already exists
	// If a session with same token exists, then don't
	// allow to create a Session
	if _, err = sStore.GetSession(token); err == nil {
		err = SessionAlreadyExistsError
		return
	}

	if session, err = NewSession(nil); err != nil {
		return
	}
	sStore.store[token] = session
	return
}

func (sStore *SessionStore) GetSession(token TokenT) (session *Session, err error) {
	var (
		isPresent bool
	)
	if session, isPresent = sStore.store[token]; !isPresent {
		err = SessionNotPresentError
	}
	return
}

func (sStore *SessionStore) RemoveSession(token TokenT) (err error) {

	sStore.storeLock.Lock()
	defer sStore.storeLock.Unlock()

	if _, err = sStore.GetSession(token); err != nil {
		return
	}

	delete(sStore.store, token)

	return
}
