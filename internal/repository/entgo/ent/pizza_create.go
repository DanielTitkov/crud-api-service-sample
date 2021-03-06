// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent/pizza"
)

// PizzaCreate is the builder for creating a Pizza entity.
type PizzaCreate struct {
	config
	mutation *PizzaMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PizzaCreate) SetCreateTime(t time.Time) *PizzaCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PizzaCreate) SetNillableCreateTime(t *time.Time) *PizzaCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PizzaCreate) SetUpdateTime(t time.Time) *PizzaCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PizzaCreate) SetNillableUpdateTime(t *time.Time) *PizzaCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PizzaCreate) SetTitle(s string) *PizzaCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PizzaCreate) SetDescription(s string) *PizzaCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PizzaCreate) SetNillableDescription(s *string) *PizzaCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetPrice sets the "price" field.
func (pc *PizzaCreate) SetPrice(i int64) *PizzaCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetDough sets the "dough" field.
func (pc *PizzaCreate) SetDough(pi pizza.Dough) *PizzaCreate {
	pc.mutation.SetDough(pi)
	return pc
}

// SetNillableDough sets the "dough" field if the given value is not nil.
func (pc *PizzaCreate) SetNillableDough(pi *pizza.Dough) *PizzaCreate {
	if pi != nil {
		pc.SetDough(*pi)
	}
	return pc
}

// Mutation returns the PizzaMutation object of the builder.
func (pc *PizzaCreate) Mutation() *PizzaMutation {
	return pc.mutation
}

// Save creates the Pizza in the database.
func (pc *PizzaCreate) Save(ctx context.Context) (*Pizza, error) {
	var (
		err  error
		node *Pizza
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PizzaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PizzaCreate) SaveX(ctx context.Context) *Pizza {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PizzaCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PizzaCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PizzaCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := pizza.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := pizza.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Dough(); !ok {
		v := pizza.DefaultDough
		pc.mutation.SetDough(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PizzaCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := pizza.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := pizza.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "price"`)}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := pizza.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Dough(); !ok {
		return &ValidationError{Name: "dough", err: errors.New(`ent: missing required field "dough"`)}
	}
	if v, ok := pc.mutation.Dough(); ok {
		if err := pizza.DoughValidator(v); err != nil {
			return &ValidationError{Name: "dough", err: fmt.Errorf(`ent: validator failed for field "dough": %w`, err)}
		}
	}
	return nil
}

func (pc *PizzaCreate) sqlSave(ctx context.Context) (*Pizza, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PizzaCreate) createSpec() (*Pizza, *sqlgraph.CreateSpec) {
	var (
		_node = &Pizza{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pizza.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pizza.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pizza.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pizza.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pizza.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pizza.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pizza.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := pc.mutation.Dough(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: pizza.FieldDough,
		})
		_node.Dough = value
	}
	return _node, _spec
}

// PizzaCreateBulk is the builder for creating many Pizza entities in bulk.
type PizzaCreateBulk struct {
	config
	builders []*PizzaCreate
}

// Save creates the Pizza entities in the database.
func (pcb *PizzaCreateBulk) Save(ctx context.Context) ([]*Pizza, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pizza, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PizzaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PizzaCreateBulk) SaveX(ctx context.Context) []*Pizza {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PizzaCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PizzaCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
