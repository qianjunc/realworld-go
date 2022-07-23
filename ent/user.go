// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/qianjunc/realworld-go/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Favorite holds the value of the favorite edge.
	Favorite []*Article `json:"favorite,omitempty"`
	// MyComments holds the value of the myComments edge.
	MyComments []*Comment `json:"myComments,omitempty"`
	// Articles holds the value of the articles edge.
	Articles []*Article `json:"articles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// FavoriteOrErr returns the Favorite value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FavoriteOrErr() ([]*Article, error) {
	if e.loadedTypes[2] {
		return e.Favorite, nil
	}
	return nil, &NotLoadedError{edge: "favorite"}
}

// MyCommentsOrErr returns the MyComments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MyCommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.MyComments, nil
	}
	return nil, &NotLoadedError{edge: "myComments"}
}

// ArticlesOrErr returns the Articles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ArticlesOrErr() ([]*Article, error) {
	if e.loadedTypes[4] {
		return e.Articles, nil
	}
	return nil, &NotLoadedError{edge: "articles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldUsername, user.FieldBio, user.FieldImage:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = value.String
			}
		case user.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				u.Image = value.String
			}
		}
	}
	return nil
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowing(u)
}

// QueryFavorite queries the "favorite" edge of the User entity.
func (u *User) QueryFavorite() *ArticleQuery {
	return (&UserClient{config: u.config}).QueryFavorite(u)
}

// QueryMyComments queries the "myComments" edge of the User entity.
func (u *User) QueryMyComments() *CommentQuery {
	return (&UserClient{config: u.config}).QueryMyComments(u)
}

// QueryArticles queries the "articles" edge of the User entity.
func (u *User) QueryArticles() *ArticleQuery {
	return (&UserClient{config: u.config}).QueryArticles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(u.Bio)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(u.Image)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}