// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/qianjunc/realworld-go/ent/article"
	"github.com/qianjunc/realworld-go/ent/comment"
	"github.com/qianjunc/realworld-go/ent/predicate"
	"github.com/qianjunc/realworld-go/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "createdAt" field.
func (cu *CommentUpdate) SetCreatedAt(s string) *CommentUpdate {
	cu.mutation.SetCreatedAt(s)
	return cu
}

// SetUpdatedAt sets the "updatedAt" field.
func (cu *CommentUpdate) SetUpdatedAt(s string) *CommentUpdate {
	cu.mutation.SetUpdatedAt(s)
	return cu
}

// SetBody sets the "body" field.
func (cu *CommentUpdate) SetBody(s string) *CommentUpdate {
	cu.mutation.SetBody(s)
	return cu
}

// SetCommentAuthorID sets the "commentAuthor" edge to the User entity by ID.
func (cu *CommentUpdate) SetCommentAuthorID(id int) *CommentUpdate {
	cu.mutation.SetCommentAuthorID(id)
	return cu
}

// SetNillableCommentAuthorID sets the "commentAuthor" edge to the User entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableCommentAuthorID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetCommentAuthorID(*id)
	}
	return cu
}

// SetCommentAuthor sets the "commentAuthor" edge to the User entity.
func (cu *CommentUpdate) SetCommentAuthor(u *User) *CommentUpdate {
	return cu.SetCommentAuthorID(u.ID)
}

// SetCommentArticleID sets the "commentArticle" edge to the Article entity by ID.
func (cu *CommentUpdate) SetCommentArticleID(id int) *CommentUpdate {
	cu.mutation.SetCommentArticleID(id)
	return cu
}

// SetNillableCommentArticleID sets the "commentArticle" edge to the Article entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableCommentArticleID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetCommentArticleID(*id)
	}
	return cu
}

// SetCommentArticle sets the "commentArticle" edge to the Article entity.
func (cu *CommentUpdate) SetCommentArticle(a *Article) *CommentUpdate {
	return cu.SetCommentArticleID(a.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearCommentAuthor clears the "commentAuthor" edge to the User entity.
func (cu *CommentUpdate) ClearCommentAuthor() *CommentUpdate {
	cu.mutation.ClearCommentAuthor()
	return cu
}

// ClearCommentArticle clears the "commentArticle" edge to the Article entity.
func (cu *CommentUpdate) ClearCommentArticle() *CommentUpdate {
	cu.mutation.ClearCommentArticle()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldBody,
		})
	}
	if cu.mutation.CommentAuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentAuthorTable,
			Columns: []string{comment.CommentAuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentAuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentAuthorTable,
			Columns: []string{comment.CommentAuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CommentArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentArticleTable,
			Columns: []string{comment.CommentArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentArticleTable,
			Columns: []string{comment.CommentArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetCreatedAt sets the "createdAt" field.
func (cuo *CommentUpdateOne) SetCreatedAt(s string) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(s)
	return cuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(s string) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(s)
	return cuo
}

// SetBody sets the "body" field.
func (cuo *CommentUpdateOne) SetBody(s string) *CommentUpdateOne {
	cuo.mutation.SetBody(s)
	return cuo
}

// SetCommentAuthorID sets the "commentAuthor" edge to the User entity by ID.
func (cuo *CommentUpdateOne) SetCommentAuthorID(id int) *CommentUpdateOne {
	cuo.mutation.SetCommentAuthorID(id)
	return cuo
}

// SetNillableCommentAuthorID sets the "commentAuthor" edge to the User entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCommentAuthorID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetCommentAuthorID(*id)
	}
	return cuo
}

// SetCommentAuthor sets the "commentAuthor" edge to the User entity.
func (cuo *CommentUpdateOne) SetCommentAuthor(u *User) *CommentUpdateOne {
	return cuo.SetCommentAuthorID(u.ID)
}

// SetCommentArticleID sets the "commentArticle" edge to the Article entity by ID.
func (cuo *CommentUpdateOne) SetCommentArticleID(id int) *CommentUpdateOne {
	cuo.mutation.SetCommentArticleID(id)
	return cuo
}

// SetNillableCommentArticleID sets the "commentArticle" edge to the Article entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCommentArticleID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetCommentArticleID(*id)
	}
	return cuo
}

// SetCommentArticle sets the "commentArticle" edge to the Article entity.
func (cuo *CommentUpdateOne) SetCommentArticle(a *Article) *CommentUpdateOne {
	return cuo.SetCommentArticleID(a.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearCommentAuthor clears the "commentAuthor" edge to the User entity.
func (cuo *CommentUpdateOne) ClearCommentAuthor() *CommentUpdateOne {
	cuo.mutation.ClearCommentAuthor()
	return cuo
}

// ClearCommentArticle clears the "commentArticle" edge to the Article entity.
func (cuo *CommentUpdateOne) ClearCommentArticle() *CommentUpdateOne {
	cuo.mutation.ClearCommentArticle()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Comment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CommentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldBody,
		})
	}
	if cuo.mutation.CommentAuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentAuthorTable,
			Columns: []string{comment.CommentAuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentAuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentAuthorTable,
			Columns: []string{comment.CommentAuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CommentArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentArticleTable,
			Columns: []string{comment.CommentArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentArticleTable,
			Columns: []string{comment.CommentArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
