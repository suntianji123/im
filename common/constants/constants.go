package constants

const (
	_ int32 = iota
	// StatusStart status
	StatusStart
	// StatusHandshake status
	StatusHandshake
	// StatusWorking status
	StatusWorking
	// StatusClosed status
	StatusClosed
)

const PacketLength = 4
const ErrorCode = -1
const SuccessCode = 0
const IllegalParam = 2

const uri = "uri"

const OffsetOfSvcId = 16
const SvcIdOfConn = 300

const Ping = SvcIdOfConn<<OffsetOfSvcId | 1
const Pong = SvcIdOfConn<<OffsetOfSvcId | 2
const LoginReq = SvcIdOfConn<<OffsetOfSvcId | 3

const LoginResp = SvcIdOfConn<<OffsetOfSvcId | 4
const KickOff = SvcIdOfConn<<OffsetOfSvcId | 7

const RespCodeServerOverload = 100

const RespCodeHeartBeatTimeout = 101

const RespCodeLoginDelay = 102

const RespCodeNotLogin = 103

const RespCodeReplaceByOther = 104
