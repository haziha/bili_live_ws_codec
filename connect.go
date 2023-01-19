package bili_live_ws_codec

import (
	"encoding/json"
	"io"
	"net/http"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

type RoomInit struct {
	Code    json.Number `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    struct {
		RoomId  json.Number `json:"room_id"`
		ShortId json.Number `json:"short_id"`
		Uid     json.Number `json:"uid"`
	} `json:"data"`
}

type DanMuInfo struct {
	Code    json.Number `json:"code"`
	Message string      `json:"message"`
	Data    struct {
		Token    string `json:"token"`
		HostList []struct {
			Host    string `json:"host"`
			Port    uint16 `json:"port"`
			WssPort uint16 `json:"wss_port"`
			WsPort  uint16 `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

func GetDanMuInfo(roomId json.Number) (danMuInfo *DanMuInfo, err error) {
	req, err := http.NewRequest("GET", "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo", nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	query := req.URL.Query()
	query.Set("id", roomId.String())
	req.URL.RawQuery = query.Encode()
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_danMuInfo := new(DanMuInfo)
	err = json.Unmarshal(body, _danMuInfo)
	if err == nil {
		danMuInfo = _danMuInfo
	}
	return
}

func GetRoomInit(roomId json.Number) (roomInit *RoomInit, err error) {
	req, err := http.NewRequest("GET", "https://api.live.bilibili.com/room/v1/Room/room_init", nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	query := req.URL.Query()
	query.Set("id", roomId.String())
	req.URL.RawQuery = query.Encode()
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_roomInit := new(RoomInit)
	err = json.Unmarshal(body, _roomInit)
	if err == nil {
		roomInit = _roomInit
	}
	return
}
