package npk

import (
    "net/http"
    "github.com/gorilla/mux"
)

func DisableCache(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        h := w.Header()
        h.Set("Cache-control", "no-cache, no-store, must-revalidate")
        h.Set("Expires", "0")
        next(w, r)
    }
}

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/", DisableCache(dashboard))
    r.HandleFunc("/login", DisableCache(login))
    r.HandleFunc("/logout", DisableCache(logout))
}
