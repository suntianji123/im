// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/im/common/data/ent/immsg"
)

// IMMsg is the model entity for the IMMsg schema.
type IMMsg struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid string `json:"sid,omitempty"`
	// FromUID holds the value of the "from_uid" field.
	FromUID int64 `json:"from_uid,omitempty"`
	// FromAppid holds the value of the "from_appid" field.
	FromAppid int `json:"from_appid,omitempty"`
	// ToUID holds the value of the "to_uid" field.
	ToUID int64 `json:"to_uid,omitempty"`
	// ToAppid holds the value of the "to_appid" field.
	ToAppid int `json:"to_appid,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel int `json:"channel,omitempty"`
	// MsgID holds the value of the "msg_id" field.
	MsgID int64 `json:"msg_id,omitempty"`
	// Cts holds the value of the "cts" field.
	Cts          int64 `json:"cts,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IMMsg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case immsg.FieldID, immsg.FieldFromUID, immsg.FieldFromAppid, immsg.FieldToUID, immsg.FieldToAppid, immsg.FieldChannel, immsg.FieldMsgID, immsg.FieldCts:
			values[i] = new(sql.NullInt64)
		case immsg.FieldSid:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IMMsg fields.
func (im *IMMsg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case immsg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			im.ID = int64(value.Int64)
		case immsg.FieldSid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value.Valid {
				im.Sid = value.String
			}
		case immsg.FieldFromUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_uid", values[i])
			} else if value.Valid {
				im.FromUID = value.Int64
			}
		case immsg.FieldFromAppid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_appid", values[i])
			} else if value.Valid {
				im.FromAppid = int(value.Int64)
			}
		case immsg.FieldToUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_uid", values[i])
			} else if value.Valid {
				im.ToUID = value.Int64
			}
		case immsg.FieldToAppid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_appid", values[i])
			} else if value.Valid {
				im.ToAppid = int(value.Int64)
			}
		case immsg.FieldChannel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				im.Channel = int(value.Int64)
			}
		case immsg.FieldMsgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field msg_id", values[i])
			} else if value.Valid {
				im.MsgID = value.Int64
			}
		case immsg.FieldCts:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cts", values[i])
			} else if value.Valid {
				im.Cts = value.Int64
			}
		default:
			im.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IMMsg.
// This includes values selected through modifiers, order, etc.
func (im *IMMsg) Value(name string) (ent.Value, error) {
	return im.selectValues.Get(name)
}

// Update returns a builder for updating this IMMsg.
// Note that you need to call IMMsg.Unwrap() before calling this method if this IMMsg
// was returned from a transaction, and the transaction was committed or rolled back.
func (im *IMMsg) Update() *IMMsgUpdateOne {
	return NewIMMsgClient(im.config).UpdateOne(im)
}

// Unwrap unwraps the IMMsg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (im *IMMsg) Unwrap() *IMMsg {
	_tx, ok := im.config.driver.(*txDriver)
	if !ok {
		panic("ent: IMMsg is not a transactional entity")
	}
	im.config.driver = _tx.drv
	return im
}

// String implements the fmt.Stringer.
func (im *IMMsg) String() string {
	var builder strings.Builder
	builder.WriteString("IMMsg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", im.ID))
	builder.WriteString("sid=")
	builder.WriteString(im.Sid)
	builder.WriteString(", ")
	builder.WriteString("from_uid=")
	builder.WriteString(fmt.Sprintf("%v", im.FromUID))
	builder.WriteString(", ")
	builder.WriteString("from_appid=")
	builder.WriteString(fmt.Sprintf("%v", im.FromAppid))
	builder.WriteString(", ")
	builder.WriteString("to_uid=")
	builder.WriteString(fmt.Sprintf("%v", im.ToUID))
	builder.WriteString(", ")
	builder.WriteString("to_appid=")
	builder.WriteString(fmt.Sprintf("%v", im.ToAppid))
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(fmt.Sprintf("%v", im.Channel))
	builder.WriteString(", ")
	builder.WriteString("msg_id=")
	builder.WriteString(fmt.Sprintf("%v", im.MsgID))
	builder.WriteString(", ")
	builder.WriteString("cts=")
	builder.WriteString(fmt.Sprintf("%v", im.Cts))
	builder.WriteByte(')')
	return builder.String()
}

// IMMsgs is a parsable slice of IMMsg.
type IMMsgs []*IMMsg
