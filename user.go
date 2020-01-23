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
{{template "base" .}}

{{define "main"}}
<form method="post">
<label>Username</label>
<input type="text" name="username"><br>
<input type="password" name="password"><br>
<input type="submit" name="Login">
</form>
{{end}}
`

func login(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        t, _ := template.New("base").Parse(BaseTemplate)
        t.New("loginForm").Parse(loginForm)
        t.ExecuteTemplate(w, "loginForm", nil)

    case "POST":
        fmt.Fprintf(w, "login post")

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
