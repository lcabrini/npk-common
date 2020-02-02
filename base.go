package npk

var BaseTemplate = `
{{define "base"}}
<!doctype html>

<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Napkoaco</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
<script defer src="https://use.fontawesome.com/releases/v5.3.1/js/all.js"></script>
</head>
<body>
{{template "navbar" .}}
<main>
{{template "main" .}}
</main>
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
</nav>
{{end}}
`
