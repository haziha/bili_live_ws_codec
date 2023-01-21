package bili_live_ws_codec

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestClient_Run(t *testing.T) {
	client := NewClient("")
	err := client.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		p := new(Packet)
		p.Heartbeat()
		for {
			err := client.WritePacket(p)
			if err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(time.Second * 15)
		}
	}()
	p := new(Packet)
	for {
		err := client.ReadPacket(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		if p.IsPvZlib() || p.IsPvBrotli() {
			fmt.Println(p.PacketHeader)
			for {
				next, err := p.DecompressNext()
				if !next || err != nil {
					break
				}
				fmt.Println(p.PacketHeader)
			}
		} else {
			if popularity, b := p.IsOpHeartbeatReply(); b {
				fmt.Println(p.PacketHeader, popularity)
			} else {
				fmt.Println(p.PacketHeader)
			}
		}
	}
}
