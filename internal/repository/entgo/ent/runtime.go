// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent/pizza"
	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pizzaMixin := schema.Pizza{}.Mixin()
	pizzaMixinFields0 := pizzaMixin[0].Fields()
	_ = pizzaMixinFields0
	pizzaFields := schema.Pizza{}.Fields()
	_ = pizzaFields
	// pizzaDescCreateTime is the schema descriptor for create_time field.
	pizzaDescCreateTime := pizzaMixinFields0[0].Descriptor()
	// pizza.DefaultCreateTime holds the default value on creation for the create_time field.
	pizza.DefaultCreateTime = pizzaDescCreateTime.Default.(func() time.Time)
	// pizzaDescUpdateTime is the schema descriptor for update_time field.
	pizzaDescUpdateTime := pizzaMixinFields0[1].Descriptor()
	// pizza.DefaultUpdateTime holds the default value on creation for the update_time field.
	pizza.DefaultUpdateTime = pizzaDescUpdateTime.Default.(func() time.Time)
	// pizza.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	pizza.UpdateDefaultUpdateTime = pizzaDescUpdateTime.UpdateDefault.(func() time.Time)
	// pizzaDescTitle is the schema descriptor for title field.
	pizzaDescTitle := pizzaFields[0].Descriptor()
	// pizza.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	pizza.TitleValidator = pizzaDescTitle.Validators[0].(func(string) error)
	// pizzaDescDescription is the schema descriptor for description field.
	pizzaDescDescription := pizzaFields[1].Descriptor()
	// pizza.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	pizza.DescriptionValidator = pizzaDescDescription.Validators[0].(func(string) error)
	// pizzaDescPrice is the schema descriptor for price field.
	pizzaDescPrice := pizzaFields[2].Descriptor()
	// pizza.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	pizza.PriceValidator = pizzaDescPrice.Validators[0].(func(int64) error)
}