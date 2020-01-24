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
<section class="hero is-fullheight is-light is-bold">
<div class="hero-body">
<div class="container">
<div class="columns is-centered">
<article class="card is-rounded">
<div class="card-content">
<h1 class="title">Login</h1>
<form method="post">
<p class="control has-icons-left">
<input type="text" class="input" name="username" placeholder="username">
<i class="fa fa-user"></i>
</p>
<p class="control has-icons-left">
<input type="password" class="input" name="password" placeholder="password">
<i class="fa fa-lock"></i>
</p>
<p class="control has-text-right">
<button class="button is-success">Login</button>
</form>
</div>
</article>
</div>
</div>
</div>
</section>
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
