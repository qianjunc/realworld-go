package testrealworld

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"testrealworld/ent"
)

// Taglist is the resolver for the taglist field.
func (r *articleResolver) Taglist(ctx context.Context, obj *ent.Article) ([]*ent.Tag, error) {
	panic(fmt.Errorf("not implemented"))
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
	return r.client.User.Create().
		SetBio("").
		SetEmail(input.Email).
		SetImage("").
		SetPassword(input.Password).
		SetToken("anything").
		SetUsername(input.Username).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input UpdateUser) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input Login) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
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
