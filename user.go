package npk

import (
    //"fmt"
    "log"
    "github.com/google/uuid"
    "net/http"
    "html/template"
)

type User struct {
    Id       uuid.UUID
    Username string
    Password string
}

func UserIdByUsername(un string) uuid.UUID {
    var uid uuid.UUID

    sql := `
    select id
    from users
    where username = $1`

    rows, err := db.Query(sql, un)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    rows.Next()
    rows.Scan(&uid)
    return uid
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
            <h1 class="title has-text-centered">Login</h1>
            <form method="post">
              <div class="field">
                <p class="control has-icon has-icons-left">
                  <input type="text" class="input" name="username" 
                      placeholder="username">
                  <span class="icon is-small is-left">
                    <i class="fa fa-user"></i>
                  </span>
                </p>
              </div>
              <div class="field">
                <p class="control has-icon has-icons-left">
                  <input type="password" class="input" name="password" 
                      placeholder="password">
                  <span class="icon is-small is-left">
                    <i class="fa fa-lock"></i>
                  </span>
                </p>
              </div>
              <div class="field">
                <p class="control has-text-right">
                  <button class="button is-success">Login</button>
                </p>
              </div>
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
    session, _ := Store.Get(r, "npk-cookie")

    switch r.Method {
    case "GET":
        t, _ := template.New("base").Parse(BaseTemplate)
        t.New("loginForm").Parse(loginForm)
        t.ExecuteTemplate(w, "loginForm", nil)

    case "POST":
        if err := r.ParseForm(); err != nil {
            log.Print("Failed to parse login form: %v", err)
            return
        }
        un := r.FormValue("username")
        pw := r.FormValue("password")
        log.Printf("Username: %s, Password: %s", un, pw)
        if authenticate(un, pw) {
            uid := UserIdByUsername(un)
            session.Values["user"] = uid
            err := session.Save(r, w)
            if err != nil {
                log.Printf("Session not saved: %v", err)
            }
            http.Redirect(w, r, "/", 301)
            //fmt.Fprintf(w, "success")
        } else {
            http.Redirect(w, r, "/login", 301)
        }

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func authenticate(un string, pw string) bool {
    var count int

    sql := `
    select count(*) 
    from users 
    where username = $1 and password = $2`

    rows, err := db.Query(sql, un, pw)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    rows.Next()
    rows.Scan(&count)
    log.Printf("User count: %d", count)
    return count == 1
}
