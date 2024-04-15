# Web aplications basics

- Handler -> like controllers in MVC-background, responsible for executing your application logic and wirte the http response headers and bodies
- Router (servermux) -> Stores a mapping between the url patterns
- Web server -> where our code will be on in the our browser

# Steps

- Creating a Handler

```go
func home(w http.ResponseWriter, r *http.Request) {
    // w -> Will be the responsible to write in the body of html
    // r -> pointer to a struct holds all information about the current request
    w.Writer([]byte("Hello")) // Will write this inside the body of html
}
```

- Router

```go
    mux := http.NewServeMux // Will create a new serve mux
    mux.HandleFunc("/", home) // Register the handler to the pattern "/"

    err := http.ListenAndServe(":4000", mux) // Will up a serve on the port 4000 with mux
```
