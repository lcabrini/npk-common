package npk

import (
    //"fmt"
    "log"
    "net/http"
    "html/template"
    // "github.com/google/uuid"
)

var DashboardTpl = `
{{template "base" .}}

{{define "main"}}
<p>User: {{.}}</p>
{{end}}
`

func dashboard(w http.ResponseWriter, r *http.Request) {
    session, err := Store.Get(r, "npk-cookie")
    if err != nil {
        log.Printf("Unable to read session: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    uid := session.Values["user"]
    t, _ := template.New("base").Parse(BaseTemplate)
    t.New("navbar").Parse(Navbar)
    t.New("dashboard").Parse(DashboardTpl)
    t.ExecuteTemplate(w, "dashboard", uid)

    /*
    if uid == nil {
        fmt.Fprint(w, "no user")
    } else {
        fmt.Fprintf(w, "user: %s", uid)
    }
    */
}
