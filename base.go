package npk

var BaseTemplate = `
{{define "base"}}
<!doctype html>

<html lang="en">
<head>
<meta charset="utf-8">
<title>Napkoaco</title>
</head>
<body>
<main>
{{template "main" .}}
</main>
</body>
</html>
{{end}}
`
