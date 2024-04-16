# Customizing HTTP headers 

Will define a route to only responds to POST method

The pattern of HTTP with good practice, lets return 405 (not allowed), to say to user that route only accept post method
To do this, we'll use the w.WriteHeader()

```go
func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST"{
    w.WriteHeader(405) // Send the 405 error to browser
    w.Write([]byte("Method not allowed"))

    return
  }
}
```


-> Nuances
- Just can call w.WriteHeader once per response, and after code has been written it can't be changed
- w.Write() by default return 200 OK status code

```go
func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST"{
    w.Header().Set("Allow", "POST") // need be 'antes' of writerHeader and write
    w.WriteHeader(405) // Send the 405 error to browser
    w.Write([]byte("Method not allowed"))
    return
  }
}
```
This header.set() will put a field show what method is allow

## Error Shortcut 

http.Error() -> call the wirteheader and write methods behind the scenes for us


## HTTP Constants

http.MethodPost -> will give the post method by default insted type "POST"
http.StatusMethodNotAllowed -> will give the status code 405 insted type 405

## Headers and content
Automatically set: Date, Content-Length and Content-Type -> By default will give the correct content type, but if not find will return 
application/octet-stream

By default a json will be sended like: text/plain, prevent this make this:

```go
w.Header().Set("Content-Type", "application/json")
```
