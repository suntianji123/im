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

const (
	_ int = iota
	ChatTypeChat
	ChatTypeGChat
)

const (
	PSyncReqNoSync   = -1
	PSyncReqNeedSync = 0
)

const (
	PushTypeNone    = 0
	PushTypeContent = 1
)

const PacketLength = 4
const ErrorCode = -1
const SuccessCode = 0
const IllegalParam = 2

const Uri = "uri"

const OffsetOfSvcId = 16
const SvcIdOfBase = 100

const SvcIdOfConn = 300
const SvcIdOfPush = 301

const SvcIdOfChat = 400

const IdOfBatch = SvcIdOfBase<<OffsetOfSvcId | 1

const Ping = SvcIdOfConn<<OffsetOfSvcId | 1
const Pong = SvcIdOfConn<<OffsetOfSvcId | 2
const LoginReq = SvcIdOfConn<<OffsetOfSvcId | 3

const LoginResp = SvcIdOfConn<<OffsetOfSvcId | 4
const KickOff = SvcIdOfConn<<OffsetOfSvcId | 7

const TransUp = SvcIdOfConn<<OffsetOfSvcId | 8
const TransDown = SvcIdOfConn<<OffsetOfSvcId | 9

// region 单聊

const ChatMsg = SvcIdOfChat<<OffsetOfSvcId | 1

const ChatMsgResp = SvcIdOfChat<<OffsetOfSvcId | 2

const SyncReq = SvcIdOfPush<<OffsetOfSvcId | 4

const RespCodeServerOverload = 100

const RespCodeHeartBeatTimeout = 101

const RespCodeLoginDelay = 102

const RespCodeNotLogin = 103

const RespCodeReplaceByOther = 104

const MqMessageSubject = "msg"
const MqSyncSubject = "sync"
