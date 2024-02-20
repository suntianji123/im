package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserInfo holds the schema definition for the UserInfo entity.
type UserInfo struct {
	ent.Schema ``
}

// Fields of the UserInfo.
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").MaxLen(100),
		field.String("password").MaxLen(100),
		field.String("nickname").MaxLen(100),
		field.String("avatar").MaxLen(300).Optional(),
		field.Int("status"),
		field.String("ext").MaxLen(500).Optional(),
	}
}

// Edges of the UserInfo.
func (UserInfo) Edges() []ent.Edge {
	return nil
}

func (UserInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("user_info"),
	}
}
