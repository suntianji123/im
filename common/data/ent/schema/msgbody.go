package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Friend holds the schema definition for the Friend entity.
type MsgBody struct {
	ent.Schema
}

// Fields of the Friend.
func (MsgBody) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("msg_id"),
		field.String("body"),
		field.Int64("cts"),
	}
}

// Edges of the Friend.
func (MsgBody) Edges() []ent.Edge {
	return nil
}

func (MsgBody) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("msg_body"),
	}
}
