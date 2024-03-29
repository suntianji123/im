// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/im/common/data/ent/userinfo"
)

// UserInfoCreate is the builder for creating a UserInfo entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uic *UserInfoCreate) SetUsername(s string) *UserInfoCreate {
	uic.mutation.SetUsername(s)
	return uic
}

// SetPassword sets the "password" field.
func (uic *UserInfoCreate) SetPassword(s string) *UserInfoCreate {
	uic.mutation.SetPassword(s)
	return uic
}

// SetNickname sets the "nickname" field.
func (uic *UserInfoCreate) SetNickname(s string) *UserInfoCreate {
	uic.mutation.SetNickname(s)
	return uic
}

// SetAvatar sets the "avatar" field.
func (uic *UserInfoCreate) SetAvatar(s string) *UserInfoCreate {
	uic.mutation.SetAvatar(s)
	return uic
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableAvatar(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetAvatar(*s)
	}
	return uic
}

// SetStatus sets the "status" field.
func (uic *UserInfoCreate) SetStatus(i int) *UserInfoCreate {
	uic.mutation.SetStatus(i)
	return uic
}

// SetExt sets the "ext" field.
func (uic *UserInfoCreate) SetExt(s string) *UserInfoCreate {
	uic.mutation.SetExt(s)
	return uic
}

// SetNillableExt sets the "ext" field if the given value is not nil.
func (uic *UserInfoCreate) SetNillableExt(s *string) *UserInfoCreate {
	if s != nil {
		uic.SetExt(*s)
	}
	return uic
}

// SetID sets the "id" field.
func (uic *UserInfoCreate) SetID(i int64) *UserInfoCreate {
	uic.mutation.SetID(i)
	return uic
}

// Mutation returns the UserInfoMutation object of the builder.
func (uic *UserInfoCreate) Mutation() *UserInfoMutation {
	return uic.mutation
}

// Save creates the UserInfo in the database.
func (uic *UserInfoCreate) Save(ctx context.Context) (*UserInfo, error) {
	return withHooks(ctx, uic.sqlSave, uic.mutation, uic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserInfoCreate) SaveX(ctx context.Context) *UserInfo {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uic *UserInfoCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserInfoCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	if _, ok := uic.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "UserInfo.username"`)}
	}
	if v, ok := uic.mutation.Username(); ok {
		if err := userinfo.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "UserInfo.username": %w`, err)}
		}
	}
	if _, ok := uic.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "UserInfo.password"`)}
	}
	if v, ok := uic.mutation.Password(); ok {
		if err := userinfo.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "UserInfo.password": %w`, err)}
		}
	}
	if _, ok := uic.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "UserInfo.nickname"`)}
	}
	if v, ok := uic.mutation.Nickname(); ok {
		if err := userinfo.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "UserInfo.nickname": %w`, err)}
		}
	}
	if v, ok := uic.mutation.Avatar(); ok {
		if err := userinfo.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "UserInfo.avatar": %w`, err)}
		}
	}
	if _, ok := uic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UserInfo.status"`)}
	}
	if v, ok := uic.mutation.Ext(); ok {
		if err := userinfo.ExtValidator(v); err != nil {
			return &ValidationError{Name: "ext", err: fmt.Errorf(`ent: validator failed for field "UserInfo.ext": %w`, err)}
		}
	}
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*UserInfo, error) {
	if err := uic.check(); err != nil {
		return nil, err
	}
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uic.mutation.id = &_node.ID
	uic.mutation.done = true
	return _node, nil
}

func (uic *UserInfoCreate) createSpec() (*UserInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserInfo{config: uic.config}
		_spec = sqlgraph.NewCreateSpec(userinfo.Table, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt64))
	)
	if id, ok := uic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uic.mutation.Username(); ok {
		_spec.SetField(userinfo.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uic.mutation.Password(); ok {
		_spec.SetField(userinfo.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uic.mutation.Nickname(); ok {
		_spec.SetField(userinfo.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uic.mutation.Avatar(); ok {
		_spec.SetField(userinfo.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uic.mutation.Status(); ok {
		_spec.SetField(userinfo.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := uic.mutation.Ext(); ok {
		_spec.SetField(userinfo.FieldExt, field.TypeString, value)
		_node.Ext = value
	}
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating many UserInfo entities in bulk.
type UserInfoCreateBulk struct {
	config
	err      error
	builders []*UserInfoCreate
}

// Save creates the UserInfo entities in the database.
func (uicb *UserInfoCreateBulk) Save(ctx context.Context) ([]*UserInfo, error) {
	if uicb.err != nil {
		return nil, uicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*UserInfo, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*UserInfo {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}
