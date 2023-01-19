package bili_live_ws_codec

import (
	"unsafe"
)

type PL = uint32 // Packet Length
type HL = uint16 // Header Length
type PV = uint16 // Protocol Version
type OP = uint32 // Operation
type SI = uint32 // Sequence ID
type BD = []byte // Body

const (
	PvNormal PV = 0 // 普通包正文不使用压缩
	PvAuth   PV = 1 // 心跳及认证包正文不使用压缩
	PvZlib   PV = 2 // 普通包正文使用zlib压缩
	PvBrotli PV = 3 // 普通包正文使用brotli压缩,解压为一个带头部的协议普通包
)

const (
	OpHeartbeat      OP = 2 // 心跳包
	OpHeartbeatReply OP = 3 //心跳包回复(人气值)
	OpNormal         OP = 5 //普通包
	OpJoinRoom       OP = 7 // 申请进入房间
	OpJoinRoomReply  OP = 8 // 进入房间回复
)

const (
	SiDefault SI = 1
)

type PacketHeader struct {
	PacketLength    PL
	HeaderLength    HL
	ProtocolVersion PV
	Operation       OP
	SequenceId      SI
}

func (pH *PacketHeader) HeaderSize() uint32 {
	var length uint32 = 0
	length += uint32(unsafe.Sizeof(pH.PacketLength))
	length += uint32(unsafe.Sizeof(pH.HeaderLength))
	length += uint32(unsafe.Sizeof(pH.ProtocolVersion))
	length += uint32(unsafe.Sizeof(pH.Operation))
	length += uint32(unsafe.Sizeof(pH.SequenceId))
	return length
}

type Packet struct {
	PacketHeader
	Body BD
}

func (_this *Packet) BodySize() uint32 {
	return uint32(len(_this.Body))
}
