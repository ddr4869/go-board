// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kafka-go/ent/caserver"
	"github.com/kafka-go/ent/predicate"
)

// CaServerDelete is the builder for deleting a CaServer entity.
type CaServerDelete struct {
	config
	hooks    []Hook
	mutation *CaServerMutation
}

// Where appends a list predicates to the CaServerDelete builder.
func (csd *CaServerDelete) Where(ps ...predicate.CaServer) *CaServerDelete {
	csd.mutation.Where(ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *CaServerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, csd.sqlExec, csd.mutation, csd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *CaServerDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *CaServerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(caserver.Table, sqlgraph.NewFieldSpec(caserver.FieldID, field.TypeInt))
	if ps := csd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	csd.mutation.done = true
	return affected, err
}

// CaServerDeleteOne is the builder for deleting a single CaServer entity.
type CaServerDeleteOne struct {
	csd *CaServerDelete
}

// Where appends a list predicates to the CaServerDelete builder.
func (csdo *CaServerDeleteOne) Where(ps ...predicate.CaServer) *CaServerDeleteOne {
	csdo.csd.mutation.Where(ps...)
	return csdo
}

// Exec executes the deletion query.
func (csdo *CaServerDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{caserver.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *CaServerDeleteOne) ExecX(ctx context.Context) {
	if err := csdo.Exec(ctx); err != nil {
		panic(err)
	}
}
