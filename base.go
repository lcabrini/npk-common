package npk

var BaseTemplate = `
{{define "base"}}
<!doctype html>

<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Napkoaco</title>
<link rel="stylesheet" href="/static/bulma/css/bulma.min.css">
<script defer src="/static/fontawesome/js/all.min.js"></script>
</head>
<body>
{{template "navbar" .}}
{{template "toolbar" .}}
<main>
{{template "main" .}}
</main>
{{template "toolbar" .}}
</body>
</html>
{{end}}
`

var Navbar = `
{{define "navbar"}}
<nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      Napkoaco
    </a>
  </div>

  <div class="navbar-end">
    <div class="navbar-item">
      <a class="button is-danger" href="/logout">
        Logout
      </a>
    </div>
  </div>
</nav>
{{end}}
`
