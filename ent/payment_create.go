// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-board/ent/payment"
	"github.com/go-board/ent/user"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetUserName sets the "user_name" field.
func (pc *PaymentCreate) SetUserName(s string) *PaymentCreate {
	pc.mutation.SetUserName(s)
	return pc
}

// SetGrade sets the "grade" field.
func (pc *PaymentCreate) SetGrade(s string) *PaymentCreate {
	pc.mutation.SetGrade(s)
	return pc
}

// SetMovieID sets the "movie_id" field.
func (pc *PaymentCreate) SetMovieID(s string) *PaymentCreate {
	pc.mutation.SetMovieID(s)
	return pc
}

// SetCreatedDate sets the "created_date" field.
func (pc *PaymentCreate) SetCreatedDate(t time.Time) *PaymentCreate {
	pc.mutation.SetCreatedDate(t)
	return pc
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableCreatedDate(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetCreatedDate(*t)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PaymentCreate) SetUserID(id string) *PaymentCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PaymentCreate) SetUser(u *User) *PaymentCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PaymentCreate) defaults() {
	if _, ok := pc.mutation.CreatedDate(); !ok {
		v := payment.DefaultCreatedDate
		pc.mutation.SetCreatedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`ent: missing required field "Payment.user_name"`)}
	}
	if _, ok := pc.mutation.Grade(); !ok {
		return &ValidationError{Name: "grade", err: errors.New(`ent: missing required field "Payment.grade"`)}
	}
	if v, ok := pc.mutation.Grade(); ok {
		if err := payment.GradeValidator(v); err != nil {
			return &ValidationError{Name: "grade", err: fmt.Errorf(`ent: validator failed for field "Payment.grade": %w`, err)}
		}
	}
	if _, ok := pc.mutation.MovieID(); !ok {
		return &ValidationError{Name: "movie_id", err: errors.New(`ent: missing required field "Payment.movie_id"`)}
	}
	if v, ok := pc.mutation.MovieID(); ok {
		if err := payment.MovieIDValidator(v); err != nil {
			return &ValidationError{Name: "movie_id", err: fmt.Errorf(`ent: validator failed for field "Payment.movie_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "created_date", err: errors.New(`ent: missing required field "Payment.created_date"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Payment.user"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.UserName(); ok {
		_spec.SetField(payment.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := pc.mutation.Grade(); ok {
		_spec.SetField(payment.FieldGrade, field.TypeString, value)
		_node.Grade = value
	}
	if value, ok := pc.mutation.MovieID(); ok {
		_spec.SetField(payment.FieldMovieID, field.TypeString, value)
		_node.MovieID = value
	}
	if value, ok := pc.mutation.CreatedDate(); ok {
		_spec.SetField(payment.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_payment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}