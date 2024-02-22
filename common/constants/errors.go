package constants

import "errors"

var (
	ErrCloseClosedSession             = errors.New("close closed session")
	ErrReceivedMsgSmallerThanExpected = errors.New("received less data than expected, EOF?")
	ErrConnectionClosed               = errors.New("client connection closed")
	ErrIllegalState                   = errors.New("error Illegal state")
)
