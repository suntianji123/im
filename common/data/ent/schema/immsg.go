package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Friend holds the schema definition for the Friend entity.
type IMMsg struct {
	ent.Schema
}

// Fields of the Friend.
func (IMMsg) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("sid"),
		field.Int64("from_uid"),
		field.Int("from_appid"),
		field.Int64("to_uid"),
		field.Int("to_appid"),
		field.Int("channel"),
		field.Int64("msg_id"),
		field.Int64("cts"),
	}
}

// Edges of the Friend.
func (IMMsg) Edges() []ent.Edge {
	return nil
}

func (IMMsg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("im_msg"),
	}
}
