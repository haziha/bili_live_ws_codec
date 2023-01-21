package bili_live_ws_codec

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestPacket_Decode(t *testing.T) {
	var uid = "0"
	var roomId = "00000000"
	var protoVer = "3"
	var platform = "web"
	var type_ = "2"
	var key = "JPOAABFZtLNZaGFAzIxka9SEg6eSgl8zn6J9WV0eJGukLpTnYehDHlNVqtMm-9a_VxeJKd-3TEhFb6Dw7ydb3Br1W64tFW5GFuKccppAD-cnchKlVTO94lh0ZRQUmKfk8IvZXa99gOYi021RCd0="

	j := fmt.Sprintf(`{"uid":%s,"roomid":%s,"protover":%s,"platform":"%s","type":%s,"key":"%s"}`, uid, roomId, protoVer, platform, type_, key)
	src := []byte{0x00, 0x00, 0x00, 0xEF, 0x00, 0x10, 0x00, 0x01, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x01}
	// \x00\x00\x00\xef\x00\x10\x00\x01\x00\x00\x00\x07\x00\x00\x00\x01
	src = append(src, []byte(j)...)
	p := Packet{}
	err := p.Decode(bytes.NewBuffer(src))
	if err != nil {
		log.Fatalln(err)
	}
	if p.PacketLength != 0x000000EF {
		log.Fatalln("packet length unequal 0x000000EF")
	}
	if p.HeaderLength != 0x0010 {
		log.Fatalln("header length unequal 0x0010")
	}
	if p.ProtocolVersion != 0x0001 {
		log.Fatalln("protocol version unequal 0x0001")
	}
	if p.Operation != 0x00000007 {
		log.Fatalln("operation unequal 0x00000007")
	}
	if p.SequenceId != 0x00000001 {
		log.Fatalln("sequence id unequal 0x00000001")
	}
	if string(p.Body) != j {
		log.Fatalln("dst body unequal src body")
	}
}
