package npk

import (
    "fmt"
    "log"
    "net/http"
    // "github.com/google/uuid"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Expires", "0")

    session, err := Store.Get(r, "npk-cookie")
    if err != nil {
        log.Printf("Unable to read session: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    uid := session.Values["user"]
    if uid == nil {
        fmt.Fprint(w, "no user")
    } else {
        fmt.Fprintf(w, "user: %s", uid)
    }
}
