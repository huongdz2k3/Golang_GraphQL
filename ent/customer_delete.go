// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer/ent/customer"
	"customer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerDelete is the builder for deleting a Customer entity.
type CustomerDelete struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerDelete builder.
func (cd *CustomerDelete) Where(ps ...predicate.Customer) *CustomerDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CustomerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CustomerMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CustomerDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CustomerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(customer.Table, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CustomerDeleteOne is the builder for deleting a single Customer entity.
type CustomerDeleteOne struct {
	cd *CustomerDelete
}

// Where appends a list predicates to the CustomerDelete builder.
func (cdo *CustomerDeleteOne) Where(ps ...predicate.Customer) *CustomerDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CustomerDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CustomerDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}