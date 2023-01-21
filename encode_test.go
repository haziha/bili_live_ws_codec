package bili_live_ws_codec

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestPacket_Encode(t *testing.T) {
	var uid = "0"
	var roomId = "00000000"
	var protoVer = "3"
	var platform = "web"
	var type_ = "2"
	var key = "JPOAABFZtLNZaGFAzIxka9SEg6eSgl8zn6J9WV0eJGukLpTnYehDHlNVqtMm-9a_VxeJKd-3TEhFb6Dw7ydb3Br1W64tFW5GFuKccppAD-cnchKlVTO94lh0ZRQUmKfk8IvZXa99gOYi021RCd0="

	j := fmt.Sprintf(`{"uid":%s,"roomid":%s,"protover":%s,"platform":"%s","type":%s,"key":"%s"}`, uid, roomId, protoVer, platform, type_, key)
	p := Packet{
		PacketHeader: PacketHeader{
			PacketLength:    0x00000000,
			HeaderLength:    0x0010,
			ProtocolVersion: 0x0001,
			Operation:       0x00000007,
			SequenceId:      0x00000001,
		},
		Body: []byte(j),
	}
	content, err := p.Encode()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Packet Encode Panic: %v", err))
	}
	src := []byte{0x00, 0x00, 0x00, 0xEF, 0x00, 0x10, 0x00, 0x01, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x01}
	// \x00\x00\x00\xef\x00\x10\x00\x01\x00\x00\x00\x07\x00\x00\x00\x01
	src = append(src, []byte(j)...)
	if !reflect.DeepEqual(src, content) {
		log.Fatalln("Packet Encode Check Fail")
	}
}
