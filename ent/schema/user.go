package schema

import (
	"testrealworld/ent/privacy"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("username").
			Unique(),
		field.String("token"),
		field.String("bio"),
		field.String("image"),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
		// favorite
		edge.From("favorite", Article.Type).
			Ref("favorited"),
		// commentAuthor
		edge.To("myComments", Comment.Type),
		// articleAuthor
		edge.To("articles", Article.Type),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Deny if not set otherwise.
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			privacy.AlwaysAllowRule(),
		},
	}
}
