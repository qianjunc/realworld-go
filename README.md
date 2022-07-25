# realword-go

## start the server
`go run server.go`

## Update schema
After updating schema.graphqls, regenerate code by running: <br/>

`go run github.com/99designs/gqlgen generate`<br/>

If there is warning about packages, just run this:<br/>

`go get github.com/99designs/gqlgen@v0.17.12`<br/>

and then rerun the first command.

if it can't generate and have the following error:

`merging type systems failed: unable to build input definition: unable to find type: github.com/qianjunc/hackernews/graph/model.NewLink 
`

try this:

`export GO111MODULE=on`

then rerun previous step
