package constants

import "errors"

var (
	ErrCloseClosedSession             = errors.New("close closed session")
	ErrReceivedMsgSmallerThanExpected = errors.New("received less data than expected, EOF?")
	ErrConnectionClosed               = errors.New("client connection closed")
	ErrIllegalState                   = errors.New("error Illegal state")
	ErrNatsNotRunning                 = errors.New("error nats not runing")
	ErrIdGeneratorWorkIdInValid       = errors.New("IdGeneartor workId Invalid")
	ErrIdGeneartorTimeCallback        = errors.New("IdGenerator time call back")
)
