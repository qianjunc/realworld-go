// Code generated by ent, DO NOT EDIT.

package article

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFavoriteCount holds the string denoting the favoritecount field in the database.
	FieldFavoriteCount = "favorite_count"
	// EdgeArticleAuthor holds the string denoting the articleauthor edge name in mutations.
	EdgeArticleAuthor = "articleAuthor"
	// EdgeFavorited holds the string denoting the favorited edge name in mutations.
	EdgeFavorited = "favorited"
	// EdgeArticleComments holds the string denoting the articlecomments edge name in mutations.
	EdgeArticleComments = "articleComments"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// Table holds the table name of the article in the database.
	Table = "articles"
	// ArticleAuthorTable is the table that holds the articleAuthor relation/edge.
	ArticleAuthorTable = "articles"
	// ArticleAuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ArticleAuthorInverseTable = "users"
	// ArticleAuthorColumn is the table column denoting the articleAuthor relation/edge.
	ArticleAuthorColumn = "user_articles"
	// FavoritedTable is the table that holds the favorited relation/edge. The primary key declared below.
	FavoritedTable = "article_favorited"
	// FavoritedInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	FavoritedInverseTable = "users"
	// ArticleCommentsTable is the table that holds the articleComments relation/edge.
	ArticleCommentsTable = "comments"
	// ArticleCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	ArticleCommentsInverseTable = "comments"
	// ArticleCommentsColumn is the table column denoting the articleComments relation/edge.
	ArticleCommentsColumn = "article_article_comments"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "article_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldSlug,
	FieldTitle,
	FieldDescription,
	FieldBody,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFavoriteCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "articles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_articles",
}

var (
	// FavoritedPrimaryKey and FavoritedColumn2 are the table columns denoting the
	// primary key for the favorited relation (M2M).
	FavoritedPrimaryKey = []string{"article_id", "user_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"article_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}