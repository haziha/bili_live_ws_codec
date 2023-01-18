package bili_live_ws_codec

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"github.com/andybalholm/brotli"
	"io"
)

func (_this *Packet) IsZlib() bool {
	return _this.ProtocolVersion == PvZlib
}

func (_this *Packet) IsBrotli() bool {
	return _this.ProtocolVersion == PvBrotli
}

func (_this *Packet) ZlibDecompress() (content []byte, err error) {
	reader, err := zlib.NewReader(bytes.NewReader(_this.Body))
	if err != nil {
		return
	}
	defer func() {
		_ = reader.Close()
	}()
	content, err = io.ReadAll(reader)
	return
}

func (_this *Packet) BrotliDecompress() (content []byte, err error) {
	reader := brotli.NewReader(bytes.NewReader(_this.Body))
	content, err = io.ReadAll(reader)
	return
}

func (_this *Packet) DeepCopy() (packet *Packet, err error) {
	_packet := new(Packet)
	_packet.PacketHeader = _this.PacketHeader
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, bytes.NewReader(_this.Body))
	if err != nil {
		return
	}
	_packet.Body = buf.Bytes()
	packet = _packet
	return
}

func (_this *Packet) Heartbeat() {
	_this.ProtocolVersion = PvAuth
	_this.Operation = OpHeartbeat
	_this.Body = []byte(`[object Object]`)
}

func (_this *Packet) JoinRoom(uid json.Number, roomId json.Number, protoVer int, platform string, type_ int, key string) (err error) {
	jr := struct {
		Uid      json.Number `json:"uid"`
		RoomId   json.Number `json:"roomid"`
		ProtoVer int         `json:"protover"`
		Platform string      `json:"platform"`
		Type     int         `json:"type"`
		Key      string      `json:"key"`
	}{Uid: uid, RoomId: roomId, ProtoVer: protoVer, Platform: platform, Type: type_, Key: key}
	body, err := json.Marshal(jr)
	if err != nil {
		return
	}
	_this.ProtocolVersion = PvAuth
	_this.Operation = OpJoinRoom
	_this.Body = body
	return
}

func (_this *Packet) JoinRoom2(roomId json.Number, key string) (err error) {
	return _this.JoinRoom("0", roomId, 3, "web", 2, key)
}
