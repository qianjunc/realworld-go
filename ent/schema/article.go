package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug"),
		field.String("title"),
		field.String("description"),
		field.String("body"),
		field.String("createdAt"),
		field.String("updatedAt"),
		field.Int("favoriteCount"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		// author
		edge.From("articleAuthor", User.Type).
			Ref("articles").
			Unique(),
		// favorite
		edge.To("favorited", User.Type),
		// comments
		edge.To("articleComments", Comment.Type),
		// tags
		edge.To("tags", Tag.Type),
	}
}