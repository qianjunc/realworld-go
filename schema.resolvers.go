package testrealworld

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"testrealworld/auth"
	"testrealworld/auth/jwt"
	"testrealworld/ent"
	"testrealworld/ent/user"

	"golang.org/x/crypto/bcrypt"
)

// Taglist is the resolver for the taglist field.
func (r *articleResolver) Taglist(ctx context.Context, obj *ent.Article) ([]*ent.Tag, error) {
	return r.client.Tag.Query().All(ctx)
}

// Favorite is the resolver for the favorite field.
func (r *articleResolver) Favorite(ctx context.Context, obj *ent.Article) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author is the resolver for the author field.
func (r *articleResolver) Author(ctx context.Context, obj *ent.Article) (*Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author is the resolver for the author field.
func (r *commentResolver) Author(ctx context.Context, obj *ent.Comment) (*Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*ent.User, error) {
	token, err := jwt.GenerateToken(input.Username)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	return r.client.User.Create().
		SetBio("").
		SetEmail(input.Email).
		SetImage("").
		SetPassword(hashedPassword).
		SetToken(token).
		SetUsername(input.Username).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input UpdateUser) (*ent.User, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil {
		return nil, fmt.Errorf("authentication failed")
	}
	users, err := r.client.User.
		Query().
		Where(
			user.Username(currentUser.Username),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return users[0].Update().
		SetBio(*input.Bio).
		SetEmail(*input.Email).
		SetImage(*input.Image).
		SetUsername(*input.Username).
		Save(ctx)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input Login) (*ent.User, error) {

	users, err := r.client.User.
		Query().
		Where(
			user.Username(*&input.Username),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	hashedPassword := users[0].Password
	if CheckPasswordHash(input.Password, hashedPassword) {
		token, err := jwt.GenerateToken(input.Username)
		if err != nil {
			return nil, err
		}
		users[0].
			Update().
			SetToken(token).
			Save(ctx)
		return users[0], nil
	}
	return nil, fmt.Errorf("Authentication failed")
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// FollowProfile is the resolver for the followProfile field.
func (r *mutationResolver) FollowProfile(ctx context.Context, username string) (*Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnfollowProfile is the resolver for the unfollowProfile field.
func (r *mutationResolver) UnfollowProfile(ctx context.Context, username string) (*Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input NewArticle) (*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, slug string, input UpdateArticle) (*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, slug string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// FavoriteArticle is the resolver for the favoriteArticle field.
func (r *mutationResolver) FavoriteArticle(ctx context.Context, slug string) (*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnfavoriteArticle is the resolver for the unfavoriteArticle field.
func (r *mutationResolver) UnfavoriteArticle(ctx context.Context, slug string) (*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, slug string, body string) (*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, slug string, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context, username string) (*Profile, error) {
	var profile Profile
	err := r.client.User.
		Query().
		Where(
			user.Username(username),
		).
		Select(user.FieldID, user.FieldUsername, user.FieldBio, user.FieldEmail, user.FieldImage).
		Scan(ctx, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// AllTags is the resolver for the allTags field.
func (r *queryResolver) AllTags(ctx context.Context) ([]*ent.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context) ([]*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesFeed is the resolver for the articlesFeed field.
func (r *queryResolver) ArticlesFeed(ctx context.Context) ([]*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, slug string) (*ent.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, slug string) ([]*ent.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article returns ArticleResolver implementation.
func (r *Resolver) Article() ArticleResolver { return &articleResolver{r} }

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type articleResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
