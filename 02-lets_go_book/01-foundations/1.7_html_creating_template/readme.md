# Render a HTML template

-> html/template

use the template.ParseFile("path") -> Will read all the template
- If you will back a directory use . like this:
```go
ts, err := template.ParseFiles("./web/html/pages/home.tmpl.html")
```
then after error handling do this:

err = ts.Execute(w, nil)
