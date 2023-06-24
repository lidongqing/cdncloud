package api

import (
	"github.com/gorilla/sessions"
)

type SessionService struct {
	Store *sessions.CookieStore
}

func NewSessionService() *SessionService {
	store := sessions.NewCookieStore([]byte("cdncloud"))
	return &SessionService{
		Store: store,
	}
}
