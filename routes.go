package npk

import (
    "github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/", dashboard)
    r.HandleFunc("/login", login)
    r.HandleFunc("/logout", logout)
}
