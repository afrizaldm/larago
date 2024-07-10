package session

import (
	"github.com/gin-contrib/sessions/cookie"
)

var Store cookie.Store = nil

func Create(secretKey string) cookie.Store {

	if Store != nil {
		return Store
	}

	Store = cookie.NewStore([]byte(secretKey))

	return Store
}
