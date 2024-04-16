# URL query strings

Add a URL a query string parameter from the user.

Need add two things:
- Get id parameter
- Convert this id to number
Use r.URL.Query().Get() will return always a string, if empty will return ""

Convert the id parameter to positive number with strconv.Atoi()

// Use fprintf() to interpolate bytes, this functions need the (writer, "string", value)

# IO.writer

Fprintf -> This accept the ResponseWriter because the https.ResponseWriter satisfies the io.writer's interface13

