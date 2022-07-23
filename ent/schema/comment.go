package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("createdAt"),
		field.String("updatedAt"),
		field.String("body"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// author
		edge.From("commentAuthor", User.Type).
			Ref("myComments").
			Unique(),
		// article
		edge.From("commentArticle", Article.Type).
			Ref("articleComments").
			Unique(),
	}
}
