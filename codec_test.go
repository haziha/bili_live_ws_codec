package bili_live_ws_codec

import (
	"fmt"
	"log"
	"testing"
)

func TestPacketHeader_HeaderSize(t *testing.T) {
	ph := PacketHeader{}
	if ph.HeaderSize() != 16 {
		log.Fatalln(fmt.Sprintf("Packet Header Length != 16: %d", ph.HeaderSize()))
	}
}
