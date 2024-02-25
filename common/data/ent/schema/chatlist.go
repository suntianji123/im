package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Friend holds the schema definition for the Friend entity.
type ChatList struct {
	ent.Schema
}

// Fields of the Friend.
func (ChatList) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("uid"),
		field.Int("channel"),
		field.Int64("chat_id"),
		field.Int64("max_msg_id"),
		field.Int64("uts"),
		field.Int("type"),
	}
}

// Edges of the Friend.
func (ChatList) Edges() []ent.Edge {
	return nil
}

func (ChatList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("chat_list"),
	}
}
