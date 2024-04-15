# Routing requests

To add news handles just create the functions of this handles and add to HandleFunc:

// News handlers

```go
func snippetView(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("something"))
}
```

// Add this handler to route

```go
    mux := http.NewServeMux // Will create a new serve mux
    mux.HandleFunc("/", home) // Register the handler to the pattern "/"
    mux.HandleFunc("/snippet/view", snippetView)
    err := http.ListenAndServe(":4000", mux) // Will up a serve on the port 4000 with mux
```

## Theory

Servemux supports two diffent types of URL: fixed paths and subtree paths:

// Will call the function just match exacly with the pattern
- fixed paths -> /snippet/view | /snippet/create

// Will call the function independently of route like /snippt/
- subtree -> /snippet/ (end with '/')


## Restricting the root url pattern
change the subtree '/' to work like a fixed path 

```go
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != '/'{
        http.NotFound(w, r)
        return 
    }
    w.Write([]byte("Hello from GoSnip"))
}
```

## DefaultServeMux

You can start a server mux just initialized your handlers, because behind the scenes these functions will be registered byte
DefaultServeMux, but the DefaultServeMux is global var, any package can access and register a route, and any third party packages

