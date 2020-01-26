package npk

import (
    "encoding/gob"
    "github.com/google/uuid"
    "github.com/gorilla/sessions"
)

var (
    key = []byte("12345678123456781234567812345678")
    Store = sessions.NewCookieStore(key)
)

func init() {
    gob.Register(&uuid.UUID{})
    Store.Options = &sessions.Options{
        Domain: "localhost",
        Path: "/",
        MaxAge: 3600 * 3,
        HttpOnly: true,
    }
}
