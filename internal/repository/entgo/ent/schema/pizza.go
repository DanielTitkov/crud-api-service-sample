package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"
)

// Pizza holds the schema definition for the Pizza entity.
type Pizza struct {
	ent.Schema
}

// Fields of the Pizza.
func (Pizza) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().NotEmpty().Immutable(),
		field.String("description").Optional().MaxLen(250),
		field.Int64("price").NonNegative(),
		field.Enum("dough").Values(
			domain.DoughThick,
			domain.DoughThin,
		).Default(domain.DoughThick),
	}
}

// Edges of the Pizza.
func (Pizza) Edges() []ent.Edge {
	return nil
}

func (Pizza) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
