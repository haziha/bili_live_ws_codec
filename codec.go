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
	PvNormal  PV = 0 // 普通数据
	PvPopular PV = 1 // 人气值
	PvZlib    PV = 2 // zlib压缩
	PvBrotli  PV = 3 // brotli压缩
)

const (
	OpJoinRoom OP = 7 // 申请进入房间
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
