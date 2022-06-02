package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Storage struct {
	ent.Schema
}

func (Storage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Immutable().Unique(),
		field.String("name"),
		field.String("location"),
	}
}

func (Storage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("storage_serial", Serial.Type),
	}
}

func (Storage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
