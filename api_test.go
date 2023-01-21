package bili_live_ws_codec

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestPacket_Heartbeat(t *testing.T) {
	p := Packet{}
	p.Heartbeat()
	_, err := p.Encode()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.PacketHeader)
	fmt.Println(string(p.Body))
}

func TestPacket_JoinRoom(t *testing.T) {
	var uid json.Number = ""
	var roomId json.Number = ""
	var protoVer = 3
	var platform = "web"
	var type_ = 2
	var key = ""

	p := Packet{}
	err := p.JoinRoom(uid, roomId, protoVer, platform, type_, key)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = p.Encode()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.PacketHeader)
	fmt.Println(string(p.Body))
}
