// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PizzasColumns holds the columns for the "pizzas" table.
	PizzasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 250},
		{Name: "price", Type: field.TypeInt64},
		{Name: "dough", Type: field.TypeEnum, Enums: []string{"thick", "thin"}, Default: "thick"},
	}
	// PizzasTable holds the schema information for the "pizzas" table.
	PizzasTable = &schema.Table{
		Name:       "pizzas",
		Columns:    PizzasColumns,
		PrimaryKey: []*schema.Column{PizzasColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PizzasTable,
	}
)

func init() {
}