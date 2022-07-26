package schema

import (
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
		field.String("username"),
		field.String("bio"),
		field.String("image"),
		field.String("password").
			Sensitive(),
		field.String("token"),
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
