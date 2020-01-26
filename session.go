package npk

import (
    "github.com/gorilla/sessions"
)

var (
    key = []byte("12345678123456781234567812345678")
    Store = sessions.NewCookieStore(key)
)
