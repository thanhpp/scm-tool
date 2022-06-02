// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thanhpp/scm/ent/item"
	"github.com/thanhpp/scm/ent/predicate"
	"github.com/thanhpp/scm/ent/serial"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *ItemUpdate) SetUpdateTime(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetSku sets the "sku" field.
func (iu *ItemUpdate) SetSku(s string) *ItemUpdate {
	iu.mutation.SetSku(s)
	return iu
}

// SetDesc sets the "desc" field.
func (iu *ItemUpdate) SetDesc(s string) *ItemUpdate {
	iu.mutation.SetDesc(s)
	return iu
}

// SetSellPrice sets the "sell_price" field.
func (iu *ItemUpdate) SetSellPrice(f float64) *ItemUpdate {
	iu.mutation.ResetSellPrice()
	iu.mutation.SetSellPrice(f)
	return iu
}

// AddSellPrice adds f to the "sell_price" field.
func (iu *ItemUpdate) AddSellPrice(f float64) *ItemUpdate {
	iu.mutation.AddSellPrice(f)
	return iu
}

// AddItemSerialIDs adds the "item_serial" edge to the Serial entity by IDs.
func (iu *ItemUpdate) AddItemSerialIDs(ids ...string) *ItemUpdate {
	iu.mutation.AddItemSerialIDs(ids...)
	return iu
}

// AddItemSerial adds the "item_serial" edges to the Serial entity.
func (iu *ItemUpdate) AddItemSerial(s ...*Serial) *ItemUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iu.AddItemSerialIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearItemSerial clears all "item_serial" edges to the Serial entity.
func (iu *ItemUpdate) ClearItemSerial() *ItemUpdate {
	iu.mutation.ClearItemSerial()
	return iu
}

// RemoveItemSerialIDs removes the "item_serial" edge to Serial entities by IDs.
func (iu *ItemUpdate) RemoveItemSerialIDs(ids ...string) *ItemUpdate {
	iu.mutation.RemoveItemSerialIDs(ids...)
	return iu
}

// RemoveItemSerial removes "item_serial" edges to Serial entities.
func (iu *ItemUpdate) RemoveItemSerial(s ...*Serial) *ItemUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iu.RemoveItemSerialIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.Sku(); ok {
		if err := item.SkuValidator(v); err != nil {
			return &ValidationError{Name: "sku", err: fmt.Errorf(`ent: validator failed for field "Item.sku": %w`, err)}
		}
	}
	if v, ok := iu.mutation.SellPrice(); ok {
		if err := item.SellPriceValidator(v); err != nil {
			return &ValidationError{Name: "sell_price", err: fmt.Errorf(`ent: validator failed for field "Item.sell_price": %w`, err)}
		}
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Sku(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldSku,
		})
	}
	if value, ok := iu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldDesc,
		})
	}
	if value, ok := iu.mutation.SellPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldSellPrice,
		})
	}
	if value, ok := iu.mutation.AddedSellPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldSellPrice,
		})
	}
	if iu.mutation.ItemSerialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedItemSerialIDs(); len(nodes) > 0 && !iu.mutation.ItemSerialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ItemSerialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *ItemUpdateOne) SetUpdateTime(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetSku sets the "sku" field.
func (iuo *ItemUpdateOne) SetSku(s string) *ItemUpdateOne {
	iuo.mutation.SetSku(s)
	return iuo
}

// SetDesc sets the "desc" field.
func (iuo *ItemUpdateOne) SetDesc(s string) *ItemUpdateOne {
	iuo.mutation.SetDesc(s)
	return iuo
}

// SetSellPrice sets the "sell_price" field.
func (iuo *ItemUpdateOne) SetSellPrice(f float64) *ItemUpdateOne {
	iuo.mutation.ResetSellPrice()
	iuo.mutation.SetSellPrice(f)
	return iuo
}

// AddSellPrice adds f to the "sell_price" field.
func (iuo *ItemUpdateOne) AddSellPrice(f float64) *ItemUpdateOne {
	iuo.mutation.AddSellPrice(f)
	return iuo
}

// AddItemSerialIDs adds the "item_serial" edge to the Serial entity by IDs.
func (iuo *ItemUpdateOne) AddItemSerialIDs(ids ...string) *ItemUpdateOne {
	iuo.mutation.AddItemSerialIDs(ids...)
	return iuo
}

// AddItemSerial adds the "item_serial" edges to the Serial entity.
func (iuo *ItemUpdateOne) AddItemSerial(s ...*Serial) *ItemUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iuo.AddItemSerialIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearItemSerial clears all "item_serial" edges to the Serial entity.
func (iuo *ItemUpdateOne) ClearItemSerial() *ItemUpdateOne {
	iuo.mutation.ClearItemSerial()
	return iuo
}

// RemoveItemSerialIDs removes the "item_serial" edge to Serial entities by IDs.
func (iuo *ItemUpdateOne) RemoveItemSerialIDs(ids ...string) *ItemUpdateOne {
	iuo.mutation.RemoveItemSerialIDs(ids...)
	return iuo
}

// RemoveItemSerial removes "item_serial" edges to Serial entities.
func (iuo *ItemUpdateOne) RemoveItemSerial(s ...*Serial) *ItemUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iuo.RemoveItemSerialIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.Sku(); ok {
		if err := item.SkuValidator(v); err != nil {
			return &ValidationError{Name: "sku", err: fmt.Errorf(`ent: validator failed for field "Item.sku": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.SellPrice(); ok {
		if err := item.SellPriceValidator(v); err != nil {
			return &ValidationError{Name: "sell_price", err: fmt.Errorf(`ent: validator failed for field "Item.sell_price": %w`, err)}
		}
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Sku(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldSku,
		})
	}
	if value, ok := iuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldDesc,
		})
	}
	if value, ok := iuo.mutation.SellPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldSellPrice,
		})
	}
	if value, ok := iuo.mutation.AddedSellPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldSellPrice,
		})
	}
	if iuo.mutation.ItemSerialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedItemSerialIDs(); len(nodes) > 0 && !iuo.mutation.ItemSerialCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ItemSerialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   item.ItemSerialTable,
			Columns: []string{item.ItemSerialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: serial.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}