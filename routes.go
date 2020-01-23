package npk

import (
    "github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/login", login)
}
