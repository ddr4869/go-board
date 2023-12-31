// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-board/ent/board"
	"github.com/go-board/ent/predicate"
	"github.com/go-board/ent/user"
)

// BoardUpdate is the builder for updating Board entities.
type BoardUpdate struct {
	config
	hooks    []Hook
	mutation *BoardMutation
}

// Where appends a list predicates to the BoardUpdate builder.
func (bu *BoardUpdate) Where(ps ...predicate.Board) *BoardUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BoardUpdate) SetTitle(s string) *BoardUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableTitle(s *string) *BoardUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetContent sets the "content" field.
func (bu *BoardUpdate) SetContent(s string) *BoardUpdate {
	bu.mutation.SetContent(s)
	return bu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableContent(s *string) *BoardUpdate {
	if s != nil {
		bu.SetContent(*s)
	}
	return bu
}

// SetWriter sets the "writer" field.
func (bu *BoardUpdate) SetWriter(s string) *BoardUpdate {
	bu.mutation.SetWriter(s)
	return bu
}

// SetNillableWriter sets the "writer" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableWriter(s *string) *BoardUpdate {
	if s != nil {
		bu.SetWriter(*s)
	}
	return bu
}

// SetView sets the "view" field.
func (bu *BoardUpdate) SetView(u uint) *BoardUpdate {
	bu.mutation.ResetView()
	bu.mutation.SetView(u)
	return bu
}

// SetNillableView sets the "view" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableView(u *uint) *BoardUpdate {
	if u != nil {
		bu.SetView(*u)
	}
	return bu
}

// AddView adds u to the "view" field.
func (bu *BoardUpdate) AddView(u int) *BoardUpdate {
	bu.mutation.AddView(u)
	return bu
}

// SetRecommend sets the "recommend" field.
func (bu *BoardUpdate) SetRecommend(u uint) *BoardUpdate {
	bu.mutation.ResetRecommend()
	bu.mutation.SetRecommend(u)
	return bu
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableRecommend(u *uint) *BoardUpdate {
	if u != nil {
		bu.SetRecommend(*u)
	}
	return bu
}

// AddRecommend adds u to the "recommend" field.
func (bu *BoardUpdate) AddRecommend(u int) *BoardUpdate {
	bu.mutation.AddRecommend(u)
	return bu
}

// SetHot sets the "hot" field.
func (bu *BoardUpdate) SetHot(b bool) *BoardUpdate {
	bu.mutation.SetHot(b)
	return bu
}

// SetNillableHot sets the "hot" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableHot(b *bool) *BoardUpdate {
	if b != nil {
		bu.SetHot(*b)
	}
	return bu
}

// SetCreatedDate sets the "created_date" field.
func (bu *BoardUpdate) SetCreatedDate(t time.Time) *BoardUpdate {
	bu.mutation.SetCreatedDate(t)
	return bu
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (bu *BoardUpdate) SetNillableCreatedDate(t *time.Time) *BoardUpdate {
	if t != nil {
		bu.SetCreatedDate(*t)
	}
	return bu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bu *BoardUpdate) SetUserID(id string) *BoardUpdate {
	bu.mutation.SetUserID(id)
	return bu
}

// SetUser sets the "user" edge to the User entity.
func (bu *BoardUpdate) SetUser(u *User) *BoardUpdate {
	return bu.SetUserID(u.ID)
}

// Mutation returns the BoardMutation object of the builder.
func (bu *BoardUpdate) Mutation() *BoardMutation {
	return bu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bu *BoardUpdate) ClearUser() *BoardUpdate {
	bu.mutation.ClearUser()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BoardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoardUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoardUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoardUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BoardUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := board.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Board.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Writer(); ok {
		if err := board.WriterValidator(v); err != nil {
			return &ValidationError{Name: "writer", err: fmt.Errorf(`ent: validator failed for field "Board.writer": %w`, err)}
		}
	}
	if _, ok := bu.mutation.UserID(); bu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Board.user"`)
	}
	return nil
}

func (bu *BoardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeUint))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Content(); ok {
		_spec.SetField(board.FieldContent, field.TypeString, value)
	}
	if value, ok := bu.mutation.Writer(); ok {
		_spec.SetField(board.FieldWriter, field.TypeString, value)
	}
	if value, ok := bu.mutation.View(); ok {
		_spec.SetField(board.FieldView, field.TypeUint, value)
	}
	if value, ok := bu.mutation.AddedView(); ok {
		_spec.AddField(board.FieldView, field.TypeUint, value)
	}
	if value, ok := bu.mutation.Recommend(); ok {
		_spec.SetField(board.FieldRecommend, field.TypeUint, value)
	}
	if value, ok := bu.mutation.AddedRecommend(); ok {
		_spec.AddField(board.FieldRecommend, field.TypeUint, value)
	}
	if value, ok := bu.mutation.Hot(); ok {
		_spec.SetField(board.FieldHot, field.TypeBool, value)
	}
	if value, ok := bu.mutation.CreatedDate(); ok {
		_spec.SetField(board.FieldCreatedDate, field.TypeTime, value)
	}
	if bu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.UserTable,
			Columns: []string{board.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.UserTable,
			Columns: []string{board.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BoardUpdateOne is the builder for updating a single Board entity.
type BoardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoardMutation
}

// SetTitle sets the "title" field.
func (buo *BoardUpdateOne) SetTitle(s string) *BoardUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableTitle(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetContent sets the "content" field.
func (buo *BoardUpdateOne) SetContent(s string) *BoardUpdateOne {
	buo.mutation.SetContent(s)
	return buo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableContent(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetContent(*s)
	}
	return buo
}

// SetWriter sets the "writer" field.
func (buo *BoardUpdateOne) SetWriter(s string) *BoardUpdateOne {
	buo.mutation.SetWriter(s)
	return buo
}

// SetNillableWriter sets the "writer" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableWriter(s *string) *BoardUpdateOne {
	if s != nil {
		buo.SetWriter(*s)
	}
	return buo
}

// SetView sets the "view" field.
func (buo *BoardUpdateOne) SetView(u uint) *BoardUpdateOne {
	buo.mutation.ResetView()
	buo.mutation.SetView(u)
	return buo
}

// SetNillableView sets the "view" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableView(u *uint) *BoardUpdateOne {
	if u != nil {
		buo.SetView(*u)
	}
	return buo
}

// AddView adds u to the "view" field.
func (buo *BoardUpdateOne) AddView(u int) *BoardUpdateOne {
	buo.mutation.AddView(u)
	return buo
}

// SetRecommend sets the "recommend" field.
func (buo *BoardUpdateOne) SetRecommend(u uint) *BoardUpdateOne {
	buo.mutation.ResetRecommend()
	buo.mutation.SetRecommend(u)
	return buo
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableRecommend(u *uint) *BoardUpdateOne {
	if u != nil {
		buo.SetRecommend(*u)
	}
	return buo
}

// AddRecommend adds u to the "recommend" field.
func (buo *BoardUpdateOne) AddRecommend(u int) *BoardUpdateOne {
	buo.mutation.AddRecommend(u)
	return buo
}

// SetHot sets the "hot" field.
func (buo *BoardUpdateOne) SetHot(b bool) *BoardUpdateOne {
	buo.mutation.SetHot(b)
	return buo
}

// SetNillableHot sets the "hot" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableHot(b *bool) *BoardUpdateOne {
	if b != nil {
		buo.SetHot(*b)
	}
	return buo
}

// SetCreatedDate sets the "created_date" field.
func (buo *BoardUpdateOne) SetCreatedDate(t time.Time) *BoardUpdateOne {
	buo.mutation.SetCreatedDate(t)
	return buo
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (buo *BoardUpdateOne) SetNillableCreatedDate(t *time.Time) *BoardUpdateOne {
	if t != nil {
		buo.SetCreatedDate(*t)
	}
	return buo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (buo *BoardUpdateOne) SetUserID(id string) *BoardUpdateOne {
	buo.mutation.SetUserID(id)
	return buo
}

// SetUser sets the "user" edge to the User entity.
func (buo *BoardUpdateOne) SetUser(u *User) *BoardUpdateOne {
	return buo.SetUserID(u.ID)
}

// Mutation returns the BoardMutation object of the builder.
func (buo *BoardUpdateOne) Mutation() *BoardMutation {
	return buo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (buo *BoardUpdateOne) ClearUser() *BoardUpdateOne {
	buo.mutation.ClearUser()
	return buo
}

// Where appends a list predicates to the BoardUpdate builder.
func (buo *BoardUpdateOne) Where(ps ...predicate.Board) *BoardUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BoardUpdateOne) Select(field string, fields ...string) *BoardUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Board entity.
func (buo *BoardUpdateOne) Save(ctx context.Context) (*Board, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoardUpdateOne) SaveX(ctx context.Context) *Board {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BoardUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoardUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BoardUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := board.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Board.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Writer(); ok {
		if err := board.WriterValidator(v); err != nil {
			return &ValidationError{Name: "writer", err: fmt.Errorf(`ent: validator failed for field "Board.writer": %w`, err)}
		}
	}
	if _, ok := buo.mutation.UserID(); buo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Board.user"`)
	}
	return nil
}

func (buo *BoardUpdateOne) sqlSave(ctx context.Context) (_node *Board, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeUint))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Board.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, board.FieldID)
		for _, f := range fields {
			if !board.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != board.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Content(); ok {
		_spec.SetField(board.FieldContent, field.TypeString, value)
	}
	if value, ok := buo.mutation.Writer(); ok {
		_spec.SetField(board.FieldWriter, field.TypeString, value)
	}
	if value, ok := buo.mutation.View(); ok {
		_spec.SetField(board.FieldView, field.TypeUint, value)
	}
	if value, ok := buo.mutation.AddedView(); ok {
		_spec.AddField(board.FieldView, field.TypeUint, value)
	}
	if value, ok := buo.mutation.Recommend(); ok {
		_spec.SetField(board.FieldRecommend, field.TypeUint, value)
	}
	if value, ok := buo.mutation.AddedRecommend(); ok {
		_spec.AddField(board.FieldRecommend, field.TypeUint, value)
	}
	if value, ok := buo.mutation.Hot(); ok {
		_spec.SetField(board.FieldHot, field.TypeBool, value)
	}
	if value, ok := buo.mutation.CreatedDate(); ok {
		_spec.SetField(board.FieldCreatedDate, field.TypeTime, value)
	}
	if buo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.UserTable,
			Columns: []string{board.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.UserTable,
			Columns: []string{board.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Board{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{board.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
