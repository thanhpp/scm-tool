package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("sku").NotEmpty().Unique(),
		field.String("desc"),
		field.Float("sell_price").Min(0),
	}
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sku"),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item_serial", Serial.Type),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

type Serial struct {
	ent.Schema
}

func (Serial) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.UUID("storage_id", uuid.New()),
		field.UUID("item_id", uuid.New()).Unique(),
	}
}

func (Serial) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "item_id", "storage_id").Unique(),
	}
}

func (Serial) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).Ref("item_serial").
			Unique().Field("item_id").Required(),
		edge.From("storage", Storage.Type).Ref("storage_serial").
			Unique().Field("storage_id").Required(),
	}
}

func (Serial) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
