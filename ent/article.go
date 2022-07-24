// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/qianjunc/realworld/ent/article"
	"github.com/qianjunc/realworld/ent/user"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt string `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt string `json:"updatedAt,omitempty"`
	// FavoriteCount holds the value of the "favoriteCount" field.
	FavoriteCount int `json:"favoriteCount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges         ArticleEdges `json:"edges"`
	user_articles *int
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// ArticleAuthor holds the value of the articleAuthor edge.
	ArticleAuthor *User `json:"articleAuthor,omitempty"`
	// Favorited holds the value of the favorited edge.
	Favorited []*User `json:"favorited,omitempty"`
	// ArticleComments holds the value of the articleComments edge.
	ArticleComments []*Comment `json:"articleComments,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]*int
}

// ArticleAuthorOrErr returns the ArticleAuthor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) ArticleAuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.ArticleAuthor == nil {
			// The edge articleAuthor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ArticleAuthor, nil
	}
	return nil, &NotLoadedError{edge: "articleAuthor"}
}

// FavoritedOrErr returns the Favorited value or an error if the edge
// was not loaded in eager-loading.
func (e ArticleEdges) FavoritedOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Favorited, nil
	}
	return nil, &NotLoadedError{edge: "favorited"}
}

// ArticleCommentsOrErr returns the ArticleComments value or an error if the edge
// was not loaded in eager-loading.
func (e ArticleEdges) ArticleCommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.ArticleComments, nil
	}
	return nil, &NotLoadedError{edge: "articleComments"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ArticleEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[3] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID, article.FieldFavoriteCount:
			values[i] = new(sql.NullInt64)
		case article.FieldSlug, article.FieldTitle, article.FieldDescription, article.FieldBody, article.FieldCreatedAt, article.FieldUpdatedAt:
			values[i] = new(sql.NullString)
		case article.ForeignKeys[0]: // user_articles
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				a.Slug = value.String
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case article.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				a.Body = value.String
			}
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				a.CreatedAt = value.String
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.String
			}
		case article.FieldFavoriteCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field favoriteCount", values[i])
			} else if value.Valid {
				a.FavoriteCount = int(value.Int64)
			}
		case article.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_articles", value)
			} else if value.Valid {
				a.user_articles = new(int)
				*a.user_articles = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryArticleAuthor queries the "articleAuthor" edge of the Article entity.
func (a *Article) QueryArticleAuthor() *UserQuery {
	return (&ArticleClient{config: a.config}).QueryArticleAuthor(a)
}

// QueryFavorited queries the "favorited" edge of the Article entity.
func (a *Article) QueryFavorited() *UserQuery {
	return (&ArticleClient{config: a.config}).QueryFavorited(a)
}

// QueryArticleComments queries the "articleComments" edge of the Article entity.
func (a *Article) QueryArticleComments() *CommentQuery {
	return (&ArticleClient{config: a.config}).QueryArticleComments(a)
}

// QueryTags queries the "tags" edge of the Article entity.
func (a *Article) QueryTags() *TagQuery {
	return (&ArticleClient{config: a.config}).QueryTags(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("slug=")
	builder.WriteString(a.Slug)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(a.Body)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(a.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(a.UpdatedAt)
	builder.WriteString(", ")
	builder.WriteString("favoriteCount=")
	builder.WriteString(fmt.Sprintf("%v", a.FavoriteCount))
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
