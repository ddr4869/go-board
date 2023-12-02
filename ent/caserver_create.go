// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kafka-go/ent/caserver"
	"github.com/kafka-go/ent/user"
)

// CaServerCreate is the builder for creating a CaServer entity.
type CaServerCreate struct {
	config
	mutation *CaServerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (csc *CaServerCreate) SetName(s string) *CaServerCreate {
	csc.mutation.SetName(s)
	return csc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (csc *CaServerCreate) SetUserID(id int) *CaServerCreate {
	csc.mutation.SetUserID(id)
	return csc
}

// SetUser sets the "user" edge to the User entity.
func (csc *CaServerCreate) SetUser(u *User) *CaServerCreate {
	return csc.SetUserID(u.ID)
}

// Mutation returns the CaServerMutation object of the builder.
func (csc *CaServerCreate) Mutation() *CaServerMutation {
	return csc.mutation
}

// Save creates the CaServer in the database.
func (csc *CaServerCreate) Save(ctx context.Context) (*CaServer, error) {
	return withHooks(ctx, csc.sqlSave, csc.mutation, csc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (csc *CaServerCreate) SaveX(ctx context.Context) *CaServer {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csc *CaServerCreate) Exec(ctx context.Context) error {
	_, err := csc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csc *CaServerCreate) ExecX(ctx context.Context) {
	if err := csc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csc *CaServerCreate) check() error {
	if _, ok := csc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CaServer.name"`)}
	}
	if v, ok := csc.mutation.Name(); ok {
		if err := caserver.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CaServer.name": %w`, err)}
		}
	}
	if _, ok := csc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "CaServer.user"`)}
	}
	return nil
}

func (csc *CaServerCreate) sqlSave(ctx context.Context) (*CaServer, error) {
	if err := csc.check(); err != nil {
		return nil, err
	}
	_node, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	csc.mutation.id = &_node.ID
	csc.mutation.done = true
	return _node, nil
}

func (csc *CaServerCreate) createSpec() (*CaServer, *sqlgraph.CreateSpec) {
	var (
		_node = &CaServer{config: csc.config}
		_spec = sqlgraph.NewCreateSpec(caserver.Table, sqlgraph.NewFieldSpec(caserver.FieldID, field.TypeInt))
	)
	if value, ok := csc.mutation.Name(); ok {
		_spec.SetField(caserver.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := csc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   caserver.UserTable,
			Columns: []string{caserver.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_caserver = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CaServerCreateBulk is the builder for creating many CaServer entities in bulk.
type CaServerCreateBulk struct {
	config
	err      error
	builders []*CaServerCreate
}

// Save creates the CaServer entities in the database.
func (cscb *CaServerCreateBulk) Save(ctx context.Context) ([]*CaServer, error) {
	if cscb.err != nil {
		return nil, cscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cscb.builders))
	nodes := make([]*CaServer, len(cscb.builders))
	mutators := make([]Mutator, len(cscb.builders))
	for i := range cscb.builders {
		func(i int, root context.Context) {
			builder := cscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CaServerMutation)
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
					_, err = mutators[i+1].Mutate(root, cscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cscb *CaServerCreateBulk) SaveX(ctx context.Context) []*CaServer {
	v, err := cscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cscb *CaServerCreateBulk) Exec(ctx context.Context) error {
	_, err := cscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscb *CaServerCreateBulk) ExecX(ctx context.Context) {
	if err := cscb.Exec(ctx); err != nil {
		panic(err)
	}
}
