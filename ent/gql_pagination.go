// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer/ent/customer"
	"customer/ent/role"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CustomerEdge is the edge representation of Customer.
type CustomerEdge struct {
	Node   *Customer `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CustomerConnection is the connection containing edges to Customer.
type CustomerConnection struct {
	Edges      []*CustomerEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *CustomerConnection) build(nodes []*Customer, pager *customerPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Customer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Customer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Customer {
			return nodes[i]
		}
	}
	c.Edges = make([]*CustomerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CustomerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CustomerPaginateOption enables pagination customization.
type CustomerPaginateOption func(*customerPager) error

// WithCustomerOrder configures pagination ordering.
func WithCustomerOrder(order *CustomerOrder) CustomerPaginateOption {
	if order == nil {
		order = DefaultCustomerOrder
	}
	o := *order
	return func(pager *customerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCustomerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCustomerFilter configures pagination filter.
func WithCustomerFilter(filter func(*CustomerQuery) (*CustomerQuery, error)) CustomerPaginateOption {
	return func(pager *customerPager) error {
		if filter == nil {
			return errors.New("CustomerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type customerPager struct {
	reverse bool
	order   *CustomerOrder
	filter  func(*CustomerQuery) (*CustomerQuery, error)
}

func newCustomerPager(opts []CustomerPaginateOption, reverse bool) (*customerPager, error) {
	pager := &customerPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCustomerOrder
	}
	return pager, nil
}

func (p *customerPager) applyFilter(query *CustomerQuery) (*CustomerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *customerPager) toCursor(c *Customer) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *customerPager) applyCursors(query *CustomerQuery, after, before *Cursor) (*CustomerQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCustomerOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *customerPager) applyOrder(query *CustomerQuery) *CustomerQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCustomerOrder.Field {
		query = query.Order(DefaultCustomerOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *customerPager) orderExpr(query *CustomerQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCustomerOrder.Field {
			b.Comma().Ident(DefaultCustomerOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Customer.
func (c *CustomerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CustomerPaginateOption,
) (*CustomerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCustomerPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CustomerConnection{Edges: []*CustomerEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CustomerOrderField defines the ordering field of Customer.
type CustomerOrderField struct {
	// Value extracts the ordering value from the given Customer.
	Value    func(*Customer) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) customer.OrderOption
	toCursor func(*Customer) Cursor
}

// CustomerOrder defines the ordering of Customer.
type CustomerOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CustomerOrderField `json:"field"`
}

// DefaultCustomerOrder is the default ordering of Customer.
var DefaultCustomerOrder = &CustomerOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CustomerOrderField{
		Value: func(c *Customer) (ent.Value, error) {
			return c.ID, nil
		},
		column: customer.FieldID,
		toTerm: customer.ByID,
		toCursor: func(c *Customer) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Customer into CustomerEdge.
func (c *Customer) ToEdge(order *CustomerOrder) *CustomerEdge {
	if order == nil {
		order = DefaultCustomerOrder
	}
	return &CustomerEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// RoleEdge is the edge representation of Role.
type RoleEdge struct {
	Node   *Role  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// RoleConnection is the connection containing edges to Role.
type RoleConnection struct {
	Edges      []*RoleEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *RoleConnection) build(nodes []*Role, pager *rolePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Role
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Role {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Role {
			return nodes[i]
		}
	}
	c.Edges = make([]*RoleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &RoleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*rolePager) error

// WithRoleOrder configures pagination ordering.
func WithRoleOrder(order *RoleOrder) RolePaginateOption {
	if order == nil {
		order = DefaultRoleOrder
	}
	o := *order
	return func(pager *rolePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRoleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRoleFilter configures pagination filter.
func WithRoleFilter(filter func(*RoleQuery) (*RoleQuery, error)) RolePaginateOption {
	return func(pager *rolePager) error {
		if filter == nil {
			return errors.New("RoleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type rolePager struct {
	reverse bool
	order   *RoleOrder
	filter  func(*RoleQuery) (*RoleQuery, error)
}

func newRolePager(opts []RolePaginateOption, reverse bool) (*rolePager, error) {
	pager := &rolePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *rolePager) applyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *rolePager) toCursor(r *Role) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *rolePager) applyCursors(query *RoleQuery, after, before *Cursor) (*RoleQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultRoleOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *rolePager) applyOrder(query *RoleQuery) *RoleQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultRoleOrder.Field {
		query = query.Order(DefaultRoleOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *rolePager) orderExpr(query *RoleQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultRoleOrder.Field {
			b.Comma().Ident(DefaultRoleOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Role.
func (r *RoleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RolePaginateOption,
) (*RoleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRolePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &RoleConnection{Edges: []*RoleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = r.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if r, err = pager.applyCursors(r, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	r = pager.applyOrder(r)
	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// RoleOrderField defines the ordering field of Role.
type RoleOrderField struct {
	// Value extracts the ordering value from the given Role.
	Value    func(*Role) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) role.OrderOption
	toCursor func(*Role) Cursor
}

// RoleOrder defines the ordering of Role.
type RoleOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *RoleOrderField `json:"field"`
}

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = &RoleOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &RoleOrderField{
		Value: func(r *Role) (ent.Value, error) {
			return r.ID, nil
		},
		column: role.FieldID,
		toTerm: role.ByID,
		toCursor: func(r *Role) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Role into RoleEdge.
func (r *Role) ToEdge(order *RoleOrder) *RoleEdge {
	if order == nil {
		order = DefaultRoleOrder
	}
	return &RoleEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}