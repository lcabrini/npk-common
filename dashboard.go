package npk

import (
    "fmt"
    "net/http"
    "github.com/google/uuid"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
    session, _ := Store.Get(r, "npk-cookie")
    uid := session.Values["user"]
    if uid := nil {
        fmt.Fprint(w, "no user")
    } else {
        fmt.Fprintf(w, "user: %s", uid)
    }
}
