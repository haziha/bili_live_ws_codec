package bili_live_ws_codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func (_this *Packet) Decode(buf *bytes.Buffer) (err error) {
	if uint32(buf.Len()) < _this.HeaderSize() {
		err = fmt.Errorf("buffer size small than header size")
		return
	}
	for _, field := range []any{&_this.PacketLength, &_this.HeaderLength, &_this.ProtocolVersion, &_this.Operation, &_this.SequenceId} {
		err = binary.Read(buf, binary.BigEndian, field)
		if err != nil {
			return
		}
	}
	if _this.PacketLength < _this.HeaderSize() {
		err = fmt.Errorf("packet length small than header size")
		return
	}
	if uint32(_this.HeaderLength) != _this.HeaderSize() {
		err = fmt.Errorf("header length unequal 16, maybe protocol updated")
		return
	}
	if uint32(buf.Len()) < _this.PacketLength-_this.HeaderSize() {
		err = fmt.Errorf("buffer size small than required body size")
		return
	}
	body := new(bytes.Buffer)
	_, err = io.CopyN(body, buf, int64(_this.PacketLength-_this.HeaderSize()))
	if err != nil {
		return
	}
	_this.Body = body.Bytes()
	return
}
