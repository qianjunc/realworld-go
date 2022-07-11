# realworld-go

## start the server
`go run server.go`

## Update schema
After updating schema.graphqls, regenerate code by running: <br/>
    `go run github.com/99designs/gqlgen generate`<br/>
If there is warning about packages, just run this:<br/>
    `go get github.com/99designs/gqlgen@v0.17.12`<br/>
and then rerun the first command.