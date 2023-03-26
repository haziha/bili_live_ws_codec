package bili_live_ws_codec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"net/url"
)

type Client struct {
	roomId     json.Number
	realRoomId json.Number
	conn       net.Conn

	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (_this *Client) RoomId() json.Number {
	return _this.roomId
}

func (_this *Client) RealRoomId() json.Number {
	return _this.realRoomId
}

func (_this *Client) Close() (err error) {
	return _this.conn.Close()
}

func (_this *Client) Connect() (err error) {
	_this.ctx, _this.cancelFunc = context.WithCancel(context.Background())
	roomInit, err := GetRoomInit(_this.roomId)
	if err != nil {
		err = fmt.Errorf("GetRoomInit call fail: %v", err)
		return
	}
	if roomInit.Code != "0" {
		err = fmt.Errorf("GetRoomInit code(%s) fail: %s", roomInit.Code, roomInit.Message)
		return
	}
	_this.realRoomId = roomInit.Data.RoomId
	danMuInfo, err := GetDanMuInfo(_this.realRoomId)
	if err != nil {
		err = fmt.Errorf("GetDanMuInfo call fail: %v", err)
		return
	}
	if danMuInfo.Code != "0" {
		err = fmt.Errorf("GetDanMuInfo code(%s) fail: %s", danMuInfo.Code, danMuInfo.Message)
		return
	}
	if len(danMuInfo.Data.HostList) == 0 {
		err = fmt.Errorf("GetDanMuInfo host list is empty")
		return
	}
	if len(danMuInfo.Data.Token) == 0 {
		err = fmt.Errorf("GetDanMuInfo token is empty")
		return
	}
	u := url.URL{
		Scheme: "wss",
		Host:   fmt.Sprintf("%s:%d", danMuInfo.Data.HostList[0].Host, danMuInfo.Data.HostList[0].WssPort),
		Path:   "sub",
	}
	_this.conn, _, _, err = ws.DefaultDialer.Dial(_this.ctx, u.String())
	if err != nil {
		return
	}
	p := new(Packet)
	err = p.JoinRoomBrotli(_this.realRoomId, danMuInfo.Data.Token)
	if err != nil {
		return
	}
	err = _this.WritePacket(p)
	return
}

func (_this *Client) ReadPacket(packet *Packet) (err error) {
	body, err := _this.Read()
	if err != nil {
		return
	}
	err = packet.Decode(bytes.NewBuffer(body))
	return
}

func (_this *Client) WritePacket(packet *Packet) (err error) {
	body, err := packet.Encode()
	if err != nil {
		return
	}
	err = _this.Write(body)
	return
}

func (_this *Client) Read() (body []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			body = nil
			err = fmt.Errorf("Client.Read panic: %v", e)
		}
	}()
	body, err = wsutil.ReadServerBinary(_this.conn)
	return
}

func (_this *Client) Write(body []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Client.Write panic: %v", e)
		}
	}()
	err = wsutil.WriteClientBinary(_this.conn, body)
	return
}

func NewClient(roomId json.Number) (cli *Client) {
	cli = new(Client)
	cli.roomId = roomId
	return cli
}
