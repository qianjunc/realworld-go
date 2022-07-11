# realworld-go

## start the server
`go run server.go`

## Update schema
After updating schema.graphqls, regenerate code by running:
    `go run github.com/99designs/gqlgen generate`
If there is warning about packages, just run this:
    `go get github.com/99designs/gqlgen@v0.17.12`
and then rerun the first command.