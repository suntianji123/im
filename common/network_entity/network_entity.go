package network_entity

import (
	"context"
	"github.com/golang/protobuf/proto"
	"net"
)

type NetworkEntity interface {
	RemoteAddr() net.Addr
	Send(ctx context.Context, message proto.Message) error
	IsActive() bool
	Close() error
}
