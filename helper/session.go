package helper

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("asshole"))
