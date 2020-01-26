package npk

import (
    "fmt"
    "net/http"
    "github.com/google/uuid"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
    session, _ := Store.Get(r, "npk-cookie")
    uid := session.Values["user"].(uuid.UUID)
    fmt.Fprintf(w, "user: %s", uid)
}
