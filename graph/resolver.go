package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/qianjunc/realworld/ent"
	"github.com/qianjunc/realworld/graph/generated"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
