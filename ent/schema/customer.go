package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name").NotEmpty(),
		field.String("address").NotEmpty(),
		field.String("license_id").NotEmpty(),
		field.String("phone_number").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty().Sensitive(),
		field.Bool("active").Default(true),
		field.Time("dob").Optional(),
		field.Int("membership_num").Nillable(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("customers").
			Unique(),
	}
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
