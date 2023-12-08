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
	"github.com/go-board/ent/payment"
	"github.com/go-board/ent/predicate"
	"github.com/go-board/ent/user"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserName sets the "user_name" field.
func (pu *PaymentUpdate) SetUserName(s string) *PaymentUpdate {
	pu.mutation.SetUserName(s)
	return pu
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableUserName(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetUserName(*s)
	}
	return pu
}

// SetGrade sets the "grade" field.
func (pu *PaymentUpdate) SetGrade(s string) *PaymentUpdate {
	pu.mutation.SetGrade(s)
	return pu
}

// SetNillableGrade sets the "grade" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableGrade(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetGrade(*s)
	}
	return pu
}

// SetMovieID sets the "movie_id" field.
func (pu *PaymentUpdate) SetMovieID(s string) *PaymentUpdate {
	pu.mutation.SetMovieID(s)
	return pu
}

// SetNillableMovieID sets the "movie_id" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableMovieID(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetMovieID(*s)
	}
	return pu
}

// SetCreatedDate sets the "created_date" field.
func (pu *PaymentUpdate) SetCreatedDate(t time.Time) *PaymentUpdate {
	pu.mutation.SetCreatedDate(t)
	return pu
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCreatedDate(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetCreatedDate(*t)
	}
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PaymentUpdate) SetUserID(id string) *PaymentUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PaymentUpdate) SetUser(u *User) *PaymentUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PaymentUpdate) ClearUser() *PaymentUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Grade(); ok {
		if err := payment.GradeValidator(v); err != nil {
			return &ValidationError{Name: "grade", err: fmt.Errorf(`ent: validator failed for field "Payment.grade": %w`, err)}
		}
	}
	if v, ok := pu.mutation.MovieID(); ok {
		if err := payment.MovieIDValidator(v); err != nil {
			return &ValidationError{Name: "movie_id", err: fmt.Errorf(`ent: validator failed for field "Payment.movie_id": %w`, err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.user"`)
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserName(); ok {
		_spec.SetField(payment.FieldUserName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Grade(); ok {
		_spec.SetField(payment.FieldGrade, field.TypeString, value)
	}
	if value, ok := pu.mutation.MovieID(); ok {
		_spec.SetField(payment.FieldMovieID, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedDate(); ok {
		_spec.SetField(payment.FieldCreatedDate, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetUserName sets the "user_name" field.
func (puo *PaymentUpdateOne) SetUserName(s string) *PaymentUpdateOne {
	puo.mutation.SetUserName(s)
	return puo
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableUserName(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetUserName(*s)
	}
	return puo
}

// SetGrade sets the "grade" field.
func (puo *PaymentUpdateOne) SetGrade(s string) *PaymentUpdateOne {
	puo.mutation.SetGrade(s)
	return puo
}

// SetNillableGrade sets the "grade" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableGrade(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetGrade(*s)
	}
	return puo
}

// SetMovieID sets the "movie_id" field.
func (puo *PaymentUpdateOne) SetMovieID(s string) *PaymentUpdateOne {
	puo.mutation.SetMovieID(s)
	return puo
}

// SetNillableMovieID sets the "movie_id" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableMovieID(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetMovieID(*s)
	}
	return puo
}

// SetCreatedDate sets the "created_date" field.
func (puo *PaymentUpdateOne) SetCreatedDate(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetCreatedDate(t)
	return puo
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCreatedDate(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetCreatedDate(*t)
	}
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PaymentUpdateOne) SetUserID(id string) *PaymentUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PaymentUpdateOne) SetUser(u *User) *PaymentUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PaymentUpdateOne) ClearUser() *PaymentUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Grade(); ok {
		if err := payment.GradeValidator(v); err != nil {
			return &ValidationError{Name: "grade", err: fmt.Errorf(`ent: validator failed for field "Payment.grade": %w`, err)}
		}
	}
	if v, ok := puo.mutation.MovieID(); ok {
		if err := payment.MovieIDValidator(v); err != nil {
			return &ValidationError{Name: "movie_id", err: fmt.Errorf(`ent: validator failed for field "Payment.movie_id": %w`, err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.user"`)
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserName(); ok {
		_spec.SetField(payment.FieldUserName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Grade(); ok {
		_spec.SetField(payment.FieldGrade, field.TypeString, value)
	}
	if value, ok := puo.mutation.MovieID(); ok {
		_spec.SetField(payment.FieldMovieID, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedDate(); ok {
		_spec.SetField(payment.FieldCreatedDate, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
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
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}