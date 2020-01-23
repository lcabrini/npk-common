package npk

import (
    "fmt"
    "github.com/google/uuid"
    "net/http"
    "html/template"
)

type User struct {
    Id       uuid.UUID
    Username string
    Password string
}

var loginForm = `
<!doctype html>

<form>
<label>Username</label>
<input type="text" name="username"><br>
<input type="password" name="password"><br>
<input type="submit" name="Login">
</form>
`

func login(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        t, _ := template.New("loginForm").Parse(loginForm)
        t.ExecuteTemplate(w, "loginForm", nil)

    case "POST":
        fmt.Fprintf(w, "login post")

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
