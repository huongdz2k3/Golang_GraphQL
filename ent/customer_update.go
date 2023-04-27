// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer/ent/customer"
	"customer/ent/predicate"
	"customer/ent/role"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetAddress sets the "address" field.
func (cu *CustomerUpdate) SetAddress(s string) *CustomerUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetLicenseID sets the "license_id" field.
func (cu *CustomerUpdate) SetLicenseID(s string) *CustomerUpdate {
	cu.mutation.SetLicenseID(s)
	return cu
}

// SetPhoneNumber sets the "phone_number" field.
func (cu *CustomerUpdate) SetPhoneNumber(s string) *CustomerUpdate {
	cu.mutation.SetPhoneNumber(s)
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetPassword sets the "password" field.
func (cu *CustomerUpdate) SetPassword(s string) *CustomerUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetActive sets the "active" field.
func (cu *CustomerUpdate) SetActive(b bool) *CustomerUpdate {
	cu.mutation.SetActive(b)
	return cu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableActive(b *bool) *CustomerUpdate {
	if b != nil {
		cu.SetActive(*b)
	}
	return cu
}

// SetDob sets the "dob" field.
func (cu *CustomerUpdate) SetDob(t time.Time) *CustomerUpdate {
	cu.mutation.SetDob(t)
	return cu
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableDob(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetDob(*t)
	}
	return cu
}

// ClearDob clears the value of the "dob" field.
func (cu *CustomerUpdate) ClearDob() *CustomerUpdate {
	cu.mutation.ClearDob()
	return cu
}

// SetMembershipNum sets the "membership_num" field.
func (cu *CustomerUpdate) SetMembershipNum(i int) *CustomerUpdate {
	cu.mutation.ResetMembershipNum()
	cu.mutation.SetMembershipNum(i)
	return cu
}

// AddMembershipNum adds i to the "membership_num" field.
func (cu *CustomerUpdate) AddMembershipNum(i int) *CustomerUpdate {
	cu.mutation.AddMembershipNum(i)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CustomerUpdate) SetCreatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCreatedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CustomerUpdate) SetUpdatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableUpdatedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (cu *CustomerUpdate) SetRolesID(id int) *CustomerUpdate {
	cu.mutation.SetRolesID(id)
	return cu
}

// SetNillableRolesID sets the "roles" edge to the Role entity by ID if the given value is not nil.
func (cu *CustomerUpdate) SetNillableRolesID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetRolesID(*id)
	}
	return cu
}

// SetRoles sets the "roles" edge to the Role entity.
func (cu *CustomerUpdate) SetRoles(r *Role) *CustomerUpdate {
	return cu.SetRolesID(r.ID)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearRoles clears the "roles" edge to the Role entity.
func (cu *CustomerUpdate) ClearRoles() *CustomerUpdate {
	cu.mutation.ClearRoles()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CustomerMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Customer.address": %w`, err)}
		}
	}
	if v, ok := cu.mutation.LicenseID(); ok {
		if err := customer.LicenseIDValidator(v); err != nil {
			return &ValidationError{Name: "license_id", err: fmt.Errorf(`ent: validator failed for field "Customer.license_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.PhoneNumber(); ok {
		if err := customer.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Customer.phone_number": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	return nil
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
	}
	if value, ok := cu.mutation.LicenseID(); ok {
		_spec.SetField(customer.FieldLicenseID, field.TypeString, value)
	}
	if value, ok := cu.mutation.PhoneNumber(); ok {
		_spec.SetField(customer.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cu.mutation.Active(); ok {
		_spec.SetField(customer.FieldActive, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Dob(); ok {
		_spec.SetField(customer.FieldDob, field.TypeTime, value)
	}
	if cu.mutation.DobCleared() {
		_spec.ClearField(customer.FieldDob, field.TypeTime)
	}
	if value, ok := cu.mutation.MembershipNum(); ok {
		_spec.SetField(customer.FieldMembershipNum, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedMembershipNum(); ok {
		_spec.AddField(customer.FieldMembershipNum, field.TypeInt, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.RolesTable,
			Columns: []string{customer.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.RolesTable,
			Columns: []string{customer.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerMutation
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CustomerUpdateOne) SetAddress(s string) *CustomerUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetLicenseID sets the "license_id" field.
func (cuo *CustomerUpdateOne) SetLicenseID(s string) *CustomerUpdateOne {
	cuo.mutation.SetLicenseID(s)
	return cuo
}

// SetPhoneNumber sets the "phone_number" field.
func (cuo *CustomerUpdateOne) SetPhoneNumber(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhoneNumber(s)
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CustomerUpdateOne) SetPassword(s string) *CustomerUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetActive sets the "active" field.
func (cuo *CustomerUpdateOne) SetActive(b bool) *CustomerUpdateOne {
	cuo.mutation.SetActive(b)
	return cuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableActive(b *bool) *CustomerUpdateOne {
	if b != nil {
		cuo.SetActive(*b)
	}
	return cuo
}

// SetDob sets the "dob" field.
func (cuo *CustomerUpdateOne) SetDob(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetDob(t)
	return cuo
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableDob(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetDob(*t)
	}
	return cuo
}

// ClearDob clears the value of the "dob" field.
func (cuo *CustomerUpdateOne) ClearDob() *CustomerUpdateOne {
	cuo.mutation.ClearDob()
	return cuo
}

// SetMembershipNum sets the "membership_num" field.
func (cuo *CustomerUpdateOne) SetMembershipNum(i int) *CustomerUpdateOne {
	cuo.mutation.ResetMembershipNum()
	cuo.mutation.SetMembershipNum(i)
	return cuo
}

// AddMembershipNum adds i to the "membership_num" field.
func (cuo *CustomerUpdateOne) AddMembershipNum(i int) *CustomerUpdateOne {
	cuo.mutation.AddMembershipNum(i)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CustomerUpdateOne) SetCreatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCreatedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableUpdatedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (cuo *CustomerUpdateOne) SetRolesID(id int) *CustomerUpdateOne {
	cuo.mutation.SetRolesID(id)
	return cuo
}

// SetNillableRolesID sets the "roles" edge to the Role entity by ID if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableRolesID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetRolesID(*id)
	}
	return cuo
}

// SetRoles sets the "roles" edge to the Role entity.
func (cuo *CustomerUpdateOne) SetRoles(r *Role) *CustomerUpdateOne {
	return cuo.SetRolesID(r.ID)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearRoles clears the "roles" edge to the Role entity.
func (cuo *CustomerUpdateOne) ClearRoles() *CustomerUpdateOne {
	cuo.mutation.ClearRoles()
	return cuo
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cuo *CustomerUpdateOne) Where(ps ...predicate.Customer) *CustomerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	return withHooks[*Customer, CustomerMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Address(); ok {
		if err := customer.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Customer.address": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.LicenseID(); ok {
		if err := customer.LicenseIDValidator(v); err != nil {
			return &ValidationError{Name: "license_id", err: fmt.Errorf(`ent: validator failed for field "Customer.license_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.PhoneNumber(); ok {
		if err := customer.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Customer.phone_number": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	return nil
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
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
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
	}
	if value, ok := cuo.mutation.LicenseID(); ok {
		_spec.SetField(customer.FieldLicenseID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.PhoneNumber(); ok {
		_spec.SetField(customer.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Active(); ok {
		_spec.SetField(customer.FieldActive, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Dob(); ok {
		_spec.SetField(customer.FieldDob, field.TypeTime, value)
	}
	if cuo.mutation.DobCleared() {
		_spec.ClearField(customer.FieldDob, field.TypeTime)
	}
	if value, ok := cuo.mutation.MembershipNum(); ok {
		_spec.SetField(customer.FieldMembershipNum, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedMembershipNum(); ok {
		_spec.AddField(customer.FieldMembershipNum, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.RolesTable,
			Columns: []string{customer.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customer.RolesTable,
			Columns: []string{customer.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
