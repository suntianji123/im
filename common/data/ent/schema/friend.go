package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Friend holds the schema definition for the Friend entity.
type Friend struct {
	ent.Schema
}

// Fields of the Friend.
func (Friend) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("uid"),
		field.Int64("peer_uid"),
		field.Int("state"),
		field.Int64("cts"),
		field.Int64("uts"),
	}
}

// Edges of the Friend.
func (Friend) Edges() []ent.Edge {
	return nil
}

func (Friend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("friend"),
	}
}
