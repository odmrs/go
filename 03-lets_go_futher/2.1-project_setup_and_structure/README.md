# Project Setup and Directory Structure

- To create a default Structure

```bash
mkdir -p bin cmd/api internal migrations remote
```
Then create the main.go and the Makefile

```
touch Makefile
touch cmd/api/main.go
```

- bin -> Compiled application binaries, ready for deployment to a production server
- cmd/api -> Directory will contain the application-specific code for the API, like the code for running the server, reading and writing HTTP requests, and managing authentication
- internal -> Ancillary packages used by the API. Like interacting with databse, data valitation, seding email...
- migrations -> Contain SQL migrations files
- remote -> Files and scripts to production server
- go.mod -> Dependencies
- Makefile -> Automating common administrative tasks (audinting our go code/ build binaries, execitomg database migrations)
