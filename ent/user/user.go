// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// EdgeFavorite holds the string denoting the favorite edge name in mutations.
	EdgeFavorite = "favorite"
	// EdgeMyComments holds the string denoting the mycomments edge name in mutations.
	EdgeMyComments = "myComments"
	// EdgeArticles holds the string denoting the articles edge name in mutations.
	EdgeArticles = "articles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FollowersTable is the table that holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_following"
	// FollowingTable is the table that holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
	// FavoriteTable is the table that holds the favorite relation/edge. The primary key declared below.
	FavoriteTable = "article_favorited"
	// FavoriteInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	FavoriteInverseTable = "articles"
	// MyCommentsTable is the table that holds the myComments relation/edge.
	MyCommentsTable = "comments"
	// MyCommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	MyCommentsInverseTable = "comments"
	// MyCommentsColumn is the table column denoting the myComments relation/edge.
	MyCommentsColumn = "user_my_comments"
	// ArticlesTable is the table that holds the articles relation/edge.
	ArticlesTable = "articles"
	// ArticlesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticlesInverseTable = "articles"
	// ArticlesColumn is the table column denoting the articles relation/edge.
	ArticlesColumn = "user_articles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldUsername,
	FieldBio,
	FieldImage,
	FieldPassword,
	FieldToken,
}

var (
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"user_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"user_id", "follower_id"}
	// FavoritePrimaryKey and FavoriteColumn2 are the table columns denoting the
	// primary key for the favorite relation (M2M).
	FavoritePrimaryKey = []string{"article_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
