package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/qianjunc/realworld/graph/generated"
	"github.com/qianjunc/realworld/graph/model"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	hashedPassword, err := HashPassword(input.Password)
	newUser, err := r.client.User.Create().
		SetBio("").
		SetEmail(input.Email).
		SetImage("").
		SetPassword(hashedPassword).
		SetUsername(input.Username).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       strconv.Itoa(newUser.ID),
		Token:    "", // add authentication later
		Email:    newUser.Email,
		Username: newUser.Username,
		Bio:      &newUser.Bio,
		Image:    &newUser.Image,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.User, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	me := users[0]
	return &model.User{
		ID:       strconv.Itoa(me.ID),
		Token:    "", // add authentication later
		Email:    me.Email,
		Username: me.Username,
		Bio:      &me.Bio,
		Image:    &me.Image,
	}, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// FollowProfile is the resolver for the followProfile field.
func (r *mutationResolver) FollowProfile(ctx context.Context, username string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnfollowProfile is the resolver for the unfollowProfile field.
func (r *mutationResolver) UnfollowProfile(ctx context.Context, username string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, slug string, input model.UpdateArticle) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, slug string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// FavoriteArticle is the resolver for the favoriteArticle field.
func (r *mutationResolver) FavoriteArticle(ctx context.Context, slug string) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnfavoriteArticle is the resolver for the unfavoriteArticle field.
func (r *mutationResolver) UnfavoriteArticle(ctx context.Context, slug string) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, slug string, body string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, slug string, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context, username string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// AllTags is the resolver for the allTags field.
func (r *queryResolver) AllTags(ctx context.Context) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesFeed is the resolver for the articlesFeed field.
func (r *queryResolver) ArticlesFeed(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, slug string) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, slug string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
