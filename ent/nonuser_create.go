// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-board/ent/nonuser"
)

// NonUserCreate is the builder for creating a NonUser entity.
type NonUserCreate struct {
	config
	mutation *NonUserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (nuc *NonUserCreate) SetName(s string) *NonUserCreate {
	nuc.mutation.SetName(s)
	return nuc
}

// SetPassword sets the "password" field.
func (nuc *NonUserCreate) SetPassword(b []byte) *NonUserCreate {
	nuc.mutation.SetPassword(b)
	return nuc
}

// SetTel sets the "tel" field.
func (nuc *NonUserCreate) SetTel(s string) *NonUserCreate {
	nuc.mutation.SetTel(s)
	return nuc
}

// SetDescription sets the "description" field.
func (nuc *NonUserCreate) SetDescription(s string) *NonUserCreate {
	nuc.mutation.SetDescription(s)
	return nuc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nuc *NonUserCreate) SetNillableDescription(s *string) *NonUserCreate {
	if s != nil {
		nuc.SetDescription(*s)
	}
	return nuc
}

// SetCreatedDate sets the "created_date" field.
func (nuc *NonUserCreate) SetCreatedDate(t time.Time) *NonUserCreate {
	nuc.mutation.SetCreatedDate(t)
	return nuc
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (nuc *NonUserCreate) SetNillableCreatedDate(t *time.Time) *NonUserCreate {
	if t != nil {
		nuc.SetCreatedDate(*t)
	}
	return nuc
}

// Mutation returns the NonUserMutation object of the builder.
func (nuc *NonUserCreate) Mutation() *NonUserMutation {
	return nuc.mutation
}

// Save creates the NonUser in the database.
func (nuc *NonUserCreate) Save(ctx context.Context) (*NonUser, error) {
	nuc.defaults()
	return withHooks(ctx, nuc.sqlSave, nuc.mutation, nuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nuc *NonUserCreate) SaveX(ctx context.Context) *NonUser {
	v, err := nuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nuc *NonUserCreate) Exec(ctx context.Context) error {
	_, err := nuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuc *NonUserCreate) ExecX(ctx context.Context) {
	if err := nuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuc *NonUserCreate) defaults() {
	if _, ok := nuc.mutation.Description(); !ok {
		v := nonuser.DefaultDescription
		nuc.mutation.SetDescription(v)
	}
	if _, ok := nuc.mutation.CreatedDate(); !ok {
		v := nonuser.DefaultCreatedDate
		nuc.mutation.SetCreatedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuc *NonUserCreate) check() error {
	if _, ok := nuc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "NonUser.name"`)}
	}
	if v, ok := nuc.mutation.Name(); ok {
		if err := nonuser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "NonUser.name": %w`, err)}
		}
	}
	if _, ok := nuc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "NonUser.password"`)}
	}
	if v, ok := nuc.mutation.Password(); ok {
		if err := nonuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "NonUser.password": %w`, err)}
		}
	}
	if _, ok := nuc.mutation.Tel(); !ok {
		return &ValidationError{Name: "tel", err: errors.New(`ent: missing required field "NonUser.tel"`)}
	}
	if v, ok := nuc.mutation.Tel(); ok {
		if err := nonuser.TelValidator(v); err != nil {
			return &ValidationError{Name: "tel", err: fmt.Errorf(`ent: validator failed for field "NonUser.tel": %w`, err)}
		}
	}
	if _, ok := nuc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "NonUser.description"`)}
	}
	if _, ok := nuc.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "created_date", err: errors.New(`ent: missing required field "NonUser.created_date"`)}
	}
	return nil
}

func (nuc *NonUserCreate) sqlSave(ctx context.Context) (*NonUser, error) {
	if err := nuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nuc.mutation.id = &_node.ID
	nuc.mutation.done = true
	return _node, nil
}

func (nuc *NonUserCreate) createSpec() (*NonUser, *sqlgraph.CreateSpec) {
	var (
		_node = &NonUser{config: nuc.config}
		_spec = sqlgraph.NewCreateSpec(nonuser.Table, sqlgraph.NewFieldSpec(nonuser.FieldID, field.TypeInt))
	)
	if value, ok := nuc.mutation.Name(); ok {
		_spec.SetField(nonuser.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nuc.mutation.Password(); ok {
		_spec.SetField(nonuser.FieldPassword, field.TypeBytes, value)
		_node.Password = value
	}
	if value, ok := nuc.mutation.Tel(); ok {
		_spec.SetField(nonuser.FieldTel, field.TypeString, value)
		_node.Tel = value
	}
	if value, ok := nuc.mutation.Description(); ok {
		_spec.SetField(nonuser.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := nuc.mutation.CreatedDate(); ok {
		_spec.SetField(nonuser.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	return _node, _spec
}

// NonUserCreateBulk is the builder for creating many NonUser entities in bulk.
type NonUserCreateBulk struct {
	config
	err      error
	builders []*NonUserCreate
}

// Save creates the NonUser entities in the database.
func (nucb *NonUserCreateBulk) Save(ctx context.Context) ([]*NonUser, error) {
	if nucb.err != nil {
		return nil, nucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nucb.builders))
	nodes := make([]*NonUser, len(nucb.builders))
	mutators := make([]Mutator, len(nucb.builders))
	for i := range nucb.builders {
		func(i int, root context.Context) {
			builder := nucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NonUserMutation)
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
					_, err = mutators[i+1].Mutate(root, nucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nucb *NonUserCreateBulk) SaveX(ctx context.Context) []*NonUser {
	v, err := nucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nucb *NonUserCreateBulk) Exec(ctx context.Context) error {
	_, err := nucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nucb *NonUserCreateBulk) ExecX(ctx context.Context) {
	if err := nucb.Exec(ctx); err != nil {
		panic(err)
	}
}
