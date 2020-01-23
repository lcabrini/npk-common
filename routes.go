package npk

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/login", login)
}

// Temporary
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Login")
}
