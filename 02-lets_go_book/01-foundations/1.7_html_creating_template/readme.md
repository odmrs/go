# Render a HTML template

-> html/template

use the template.ParseFile("path") -> Will read all the template
- If you will back a directory use . like this:
```go
ts, err := template.ParseFiles("./web/html/pages/home.tmpl.html")
```
then after error handling do this:

err = ts.Execute(w, nil)


## Add a template base to use in all pages

will be like that:

```html
{{define "base"}}
<!DOCTYPE>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{template "title" .}} - Snippetbox</title>
  </head>
  <body>
    <header>
      <h1><a href="/">GoSnip</a></h1>
    </header>
    <main>
      {{template "main"}}
    </main>
    <footer>
      Powered by <a href="https://www.github.com/odmrs">Marcos (odmrs)</a>
    </footer>
  </body>
</html>
{end}
```

{{define "name"}} - {{end}} - define a template called name

{{template "title" .}} and {{template "main" .}} will invoke other named templates called title and main
- The dot inside this templating represents any dynamic data


in anothers html you will use need be like that:

File: ../pages/home.tmpl.html

{{define "tittle"}}Home{{end}}

{{define "main"}}
  <h2>Latest Snippet</h2>
  <p>The's nothing to see here yet!!</p>
{{end}}


In you code put in a slice string two path to your files, first you base template then your page:

```go
  files := []string{
    "./web/html/base/base.tmpl.html",
    "./web/html/pages/home.tmpl.html",
  }
	// using ttemplate.parsefiles() -> will read a template
	ts, err := template.ParseFiles(files ...)
	
```
