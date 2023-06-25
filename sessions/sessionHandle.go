package sessions

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
)

type SessionHandle struct {
	Store *RedisStore
}

func NewSessionHandle() *SessionHandle {
	rdCmd := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	store, err := NewRedisStore(rdCmd, []byte("secret"))
	if err != nil {
		panic(err)
	}
	store.SetMaxAge(10 * 24 * 3600)
	return &SessionHandle{
		Store: store,
	}
}

func (s *SessionHandle) SetSession(ctx context.Context, key string, value string) (bool, error) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		if ht, ok := tr.(*http.Transport); ok {
			// get a session
			session, err := s.Store.Get(ht, "session")
			if err != nil {
				return false, errors.InternalServer("INTERNAL_ERROR", "get session error")
			}

			// modified the value of the key, and save
			session.Values[key] = value
			if err = session.Save(ht); err != nil {
				return false, errors.InternalServer("INTERNAL_ERROR", "save session error")
			}

			// delete session
			//session.Options.MaxAge = -1
			//if err = session.Save(ht); err != nil {
			//	return nil, errors.InternalServer("INTERNAL_ERROR", "save session error")
			//}
		}
	}
	return true, nil
}

func (s *SessionHandle) GetSession(ctx context.Context, key string) (string, error) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		if ht, ok := tr.(*http.Transport); ok {
			// get a session
			session, err := s.Store.Get(ht, "session")
			if err != nil {
				return "", errors.InternalServer("INTERNAL_ERROR", "get session error")
			}

			// get the value of the key
			value := session.Values[key]
			if value == nil {
				return "", nil
			}
			return value.(string), nil
		}
	}
	return "", nil
}
