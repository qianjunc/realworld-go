// Code generated by ent, DO NOT EDIT.

package comment

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// EdgeCommentAuthor holds the string denoting the commentauthor edge name in mutations.
	EdgeCommentAuthor = "commentAuthor"
	// EdgeCommentArticle holds the string denoting the commentarticle edge name in mutations.
	EdgeCommentArticle = "commentArticle"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// CommentAuthorTable is the table that holds the commentAuthor relation/edge.
	CommentAuthorTable = "comments"
	// CommentAuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CommentAuthorInverseTable = "users"
	// CommentAuthorColumn is the table column denoting the commentAuthor relation/edge.
	CommentAuthorColumn = "user_my_comments"
	// CommentArticleTable is the table that holds the commentArticle relation/edge.
	CommentArticleTable = "comments"
	// CommentArticleInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	CommentArticleInverseTable = "articles"
	// CommentArticleColumn is the table column denoting the commentArticle relation/edge.
	CommentArticleColumn = "article_article_comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBody,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"article_article_comments",
	"user_my_comments",
}

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