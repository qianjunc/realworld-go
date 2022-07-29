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

## branches
There are two versions of this project
### try-ent branch
- does not connect to ent schema
- use ent as a normal orm: directly call its function to manipulate and access data
- then put them into model generated from graphql schema object

### try-connect-model branch
- connect to ent schema
- return ent schema type and use ent field collection features to access different object
