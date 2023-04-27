// Code generated by ent, DO NOT EDIT.

package role

import (
	"customer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldID, id))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v Name) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v Name) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...Name) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...Name) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldName, vs...))
}

// HasCustomers applies the HasEdge predicate on the "customers" edge.
func HasCustomers() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CustomersTable, CustomersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomersWith applies the HasEdge predicate on the "customers" edge with a given conditions (other predicates).
func HasCustomersWith(preds ...predicate.Customer) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newCustomersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		p(s.Not())
	})
}
