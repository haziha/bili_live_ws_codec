package bili_live_ws_codec

import (
	"fmt"
	"log"
	"testing"
)

func TestGetRoomInit(t *testing.T) {
	info, err := GetRoomInit("")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(info)
}

func TestGetDanMuInfo(t *testing.T) {
	info, err := GetDanMuInfo("")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(info)
}
