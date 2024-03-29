// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/im/common/data/ent/friend"
	"github.com/im/common/data/ent/predicate"
)

// FriendUpdate is the builder for updating Friend entities.
type FriendUpdate struct {
	config
	hooks    []Hook
	mutation *FriendMutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fu *FriendUpdate) Where(ps ...predicate.Friend) *FriendUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUID sets the "uid" field.
func (fu *FriendUpdate) SetUID(i int64) *FriendUpdate {
	fu.mutation.ResetUID()
	fu.mutation.SetUID(i)
	return fu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableUID(i *int64) *FriendUpdate {
	if i != nil {
		fu.SetUID(*i)
	}
	return fu
}

// AddUID adds i to the "uid" field.
func (fu *FriendUpdate) AddUID(i int64) *FriendUpdate {
	fu.mutation.AddUID(i)
	return fu
}

// SetPeerUID sets the "peer_uid" field.
func (fu *FriendUpdate) SetPeerUID(i int64) *FriendUpdate {
	fu.mutation.ResetPeerUID()
	fu.mutation.SetPeerUID(i)
	return fu
}

// SetNillablePeerUID sets the "peer_uid" field if the given value is not nil.
func (fu *FriendUpdate) SetNillablePeerUID(i *int64) *FriendUpdate {
	if i != nil {
		fu.SetPeerUID(*i)
	}
	return fu
}

// AddPeerUID adds i to the "peer_uid" field.
func (fu *FriendUpdate) AddPeerUID(i int64) *FriendUpdate {
	fu.mutation.AddPeerUID(i)
	return fu
}

// SetState sets the "state" field.
func (fu *FriendUpdate) SetState(i int) *FriendUpdate {
	fu.mutation.ResetState()
	fu.mutation.SetState(i)
	return fu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableState(i *int) *FriendUpdate {
	if i != nil {
		fu.SetState(*i)
	}
	return fu
}

// AddState adds i to the "state" field.
func (fu *FriendUpdate) AddState(i int) *FriendUpdate {
	fu.mutation.AddState(i)
	return fu
}

// SetCts sets the "cts" field.
func (fu *FriendUpdate) SetCts(i int64) *FriendUpdate {
	fu.mutation.ResetCts()
	fu.mutation.SetCts(i)
	return fu
}

// SetNillableCts sets the "cts" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableCts(i *int64) *FriendUpdate {
	if i != nil {
		fu.SetCts(*i)
	}
	return fu
}

// AddCts adds i to the "cts" field.
func (fu *FriendUpdate) AddCts(i int64) *FriendUpdate {
	fu.mutation.AddCts(i)
	return fu
}

// SetUts sets the "uts" field.
func (fu *FriendUpdate) SetUts(i int64) *FriendUpdate {
	fu.mutation.ResetUts()
	fu.mutation.SetUts(i)
	return fu
}

// SetNillableUts sets the "uts" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableUts(i *int64) *FriendUpdate {
	if i != nil {
		fu.SetUts(*i)
	}
	return fu
}

// AddUts adds i to the "uts" field.
func (fu *FriendUpdate) AddUts(i int64) *FriendUpdate {
	fu.mutation.AddUts(i)
	return fu
}

// Mutation returns the FriendMutation object of the builder.
func (fu *FriendUpdate) Mutation() *FriendMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FriendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeInt64))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UID(); ok {
		_spec.SetField(friend.FieldUID, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedUID(); ok {
		_spec.AddField(friend.FieldUID, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.PeerUID(); ok {
		_spec.SetField(friend.FieldPeerUID, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedPeerUID(); ok {
		_spec.AddField(friend.FieldPeerUID, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.State(); ok {
		_spec.SetField(friend.FieldState, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedState(); ok {
		_spec.AddField(friend.FieldState, field.TypeInt, value)
	}
	if value, ok := fu.mutation.Cts(); ok {
		_spec.SetField(friend.FieldCts, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedCts(); ok {
		_spec.AddField(friend.FieldCts, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.Uts(); ok {
		_spec.SetField(friend.FieldUts, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedUts(); ok {
		_spec.AddField(friend.FieldUts, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendUpdateOne is the builder for updating a single Friend entity.
type FriendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendMutation
}

// SetUID sets the "uid" field.
func (fuo *FriendUpdateOne) SetUID(i int64) *FriendUpdateOne {
	fuo.mutation.ResetUID()
	fuo.mutation.SetUID(i)
	return fuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableUID(i *int64) *FriendUpdateOne {
	if i != nil {
		fuo.SetUID(*i)
	}
	return fuo
}

// AddUID adds i to the "uid" field.
func (fuo *FriendUpdateOne) AddUID(i int64) *FriendUpdateOne {
	fuo.mutation.AddUID(i)
	return fuo
}

// SetPeerUID sets the "peer_uid" field.
func (fuo *FriendUpdateOne) SetPeerUID(i int64) *FriendUpdateOne {
	fuo.mutation.ResetPeerUID()
	fuo.mutation.SetPeerUID(i)
	return fuo
}

// SetNillablePeerUID sets the "peer_uid" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillablePeerUID(i *int64) *FriendUpdateOne {
	if i != nil {
		fuo.SetPeerUID(*i)
	}
	return fuo
}

// AddPeerUID adds i to the "peer_uid" field.
func (fuo *FriendUpdateOne) AddPeerUID(i int64) *FriendUpdateOne {
	fuo.mutation.AddPeerUID(i)
	return fuo
}

// SetState sets the "state" field.
func (fuo *FriendUpdateOne) SetState(i int) *FriendUpdateOne {
	fuo.mutation.ResetState()
	fuo.mutation.SetState(i)
	return fuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableState(i *int) *FriendUpdateOne {
	if i != nil {
		fuo.SetState(*i)
	}
	return fuo
}

// AddState adds i to the "state" field.
func (fuo *FriendUpdateOne) AddState(i int) *FriendUpdateOne {
	fuo.mutation.AddState(i)
	return fuo
}

// SetCts sets the "cts" field.
func (fuo *FriendUpdateOne) SetCts(i int64) *FriendUpdateOne {
	fuo.mutation.ResetCts()
	fuo.mutation.SetCts(i)
	return fuo
}

// SetNillableCts sets the "cts" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableCts(i *int64) *FriendUpdateOne {
	if i != nil {
		fuo.SetCts(*i)
	}
	return fuo
}

// AddCts adds i to the "cts" field.
func (fuo *FriendUpdateOne) AddCts(i int64) *FriendUpdateOne {
	fuo.mutation.AddCts(i)
	return fuo
}

// SetUts sets the "uts" field.
func (fuo *FriendUpdateOne) SetUts(i int64) *FriendUpdateOne {
	fuo.mutation.ResetUts()
	fuo.mutation.SetUts(i)
	return fuo
}

// SetNillableUts sets the "uts" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableUts(i *int64) *FriendUpdateOne {
	if i != nil {
		fuo.SetUts(*i)
	}
	return fuo
}

// AddUts adds i to the "uts" field.
func (fuo *FriendUpdateOne) AddUts(i int64) *FriendUpdateOne {
	fuo.mutation.AddUts(i)
	return fuo
}

// Mutation returns the FriendMutation object of the builder.
func (fuo *FriendUpdateOne) Mutation() *FriendMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fuo *FriendUpdateOne) Where(ps ...predicate.Friend) *FriendUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendUpdateOne) Select(field string, fields ...string) *FriendUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friend entity.
func (fuo *FriendUpdateOne) Save(ctx context.Context) (*Friend, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendUpdateOne) SaveX(ctx context.Context) *Friend {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FriendUpdateOne) sqlSave(ctx context.Context) (_node *Friend, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeInt64))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friend.FieldID)
		for _, f := range fields {
			if !friend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UID(); ok {
		_spec.SetField(friend.FieldUID, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedUID(); ok {
		_spec.AddField(friend.FieldUID, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.PeerUID(); ok {
		_spec.SetField(friend.FieldPeerUID, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedPeerUID(); ok {
		_spec.AddField(friend.FieldPeerUID, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.State(); ok {
		_spec.SetField(friend.FieldState, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedState(); ok {
		_spec.AddField(friend.FieldState, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.Cts(); ok {
		_spec.SetField(friend.FieldCts, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedCts(); ok {
		_spec.AddField(friend.FieldCts, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.Uts(); ok {
		_spec.SetField(friend.FieldUts, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedUts(); ok {
		_spec.AddField(friend.FieldUts, field.TypeInt64, value)
	}
	_node = &Friend{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
