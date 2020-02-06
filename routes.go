package npk

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    //"github.com/gorilla/sessions"
)

func DisableCache(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        h := w.Header()
        h.Set("Cache-control", "no-cache, no-store, must-revalidate")
        h.Set("Expires", "0")
        next(w, r)
    }
}

func Authenticated(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, err := Store.Get(r, "npk-cookie")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        uid := session.Values["user"]
        if uid == nil {
            session.AddFlash("Authentication required.")
            err = session.Save(r, w)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            http.Redirect(w, r, "/login", http.StatusFound)
        }

        next(w, r)
    }
}

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/", DisableCache(Authenticated(dashboard)))
    r.HandleFunc("/login", DisableCache(login))
    r.HandleFunc("/logout", DisableCache(Authenticated(logout)))

    r.HandleFunc("/bulma", func(w http.ResponseWriter, r *http.Request) {
        h := w.Header()
        h.Set("Content-Type", "text/css")
        fmt.Fprintf(w, Bulma)
    })
    r.HandleFunc("/fa", func(w http.ResponseWriter, r *http.Request) {
        h := w.Header()
        h.Set("Content-Type", "application/javascript")
        fmt.Fprintf(w, FontAwesome)
    })
}
