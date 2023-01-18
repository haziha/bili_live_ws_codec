package bili_live_ws_codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (_this *Packet) Encode() (content BD, err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("packet encode panic: %v", err1)
		}
	}()
	_this.PacketLength = _this.HeaderSize() + _this.BodySize()
	_this.HeaderLength = uint16(_this.HeaderSize())
	_this.SequenceId = SiDefault
	buf := bytes.Buffer{}
	for _, field := range []any{_this.PacketLength, _this.HeaderLength, _this.ProtocolVersion, _this.Operation, _this.SequenceId} {
		err = binary.Write(&buf, binary.BigEndian, field)
		if err != nil {
			return
		}
	}
	_, err = buf.Write(_this.Body)
	if err != nil {
		return
	}
	content = buf.Bytes()
	return
}
