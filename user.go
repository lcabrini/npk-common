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
<div class="container columns">
<div class="card column is-half is-offset-one-quarter">
<header class="card-header">
<p class="card-header-title">Login</p>
</header>
<div class="card-content">
<form method="post">
<div class="field">
<label class="label">Username</label>
<div class="control has-icon-left">
<input type="text" class="input" name="username" placeholder="username">
<span class="icon is-small is-left">
<i class="fa fa-user"></i>
</span>
</div>
<div class="field">
<label class="label">Password</label>
<div class="control has-icon-left">
<input type="password" class="input" name="password" placeholder="password">
<span class="icon is-small is-left">
<i class="fa fa-lock"></i>
</span>
</div>
</div>
<div class="field has-text-right">
<button class="button is-success">Login</button>
</div>
</div>
</form>
</div>
</div>
</div>
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
