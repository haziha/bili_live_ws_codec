package bili_live_ws_codec

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"github.com/andybalholm/brotli"
	"io"
)

func (_this *Packet) IsPvZlib() bool {
	return _this.ProtocolVersion == PvZlib
}

func (_this *Packet) IsPvBrotli() bool {
	return _this.ProtocolVersion == PvBrotli
}

func (_this *Packet) IsPvNormal() bool {
	return _this.ProtocolVersion == PvNormal
}

func (_this *Packet) IsPvAuth() bool {
	return _this.ProtocolVersion == PvAuth
}

func (_this *Packet) IsOpHeartbeatReply() (popularity uint32, b bool) {
	b = _this.Operation == OpHeartbeatReply
	if b {
		err := binary.Read(bytes.NewReader(_this.Body), binary.BigEndian, &popularity)
		if err != nil {
			popularity = 0
		}
	}
	return
}

func (_this *Packet) IsOpNormal() bool {
	return _this.Operation == OpNormal
}

func (_this *Packet) IsOpJoinRoomReply() bool {
	return _this.Operation == OpJoinRoomReply
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

func (_this *Packet) JoinRoomBrotli(roomId json.Number, key string) (err error) {
	return _this.JoinRoom("0", roomId, 3, "web", 2, key)
}

func (_this *Packet) JoinRoomZlib(roomId json.Number, key string) (err error) {
	return _this.JoinRoom("0", roomId, 2, "web", 2, key)
}

func (_this *Packet) DecompressNext() (next bool, err error) {
	if _this.bodyBuffer == nil {
		_this.bodyBuffer = new(bytes.Buffer)
	}
	if _this.IsPvZlib() {
		content, err := _this.ZlibDecompress()
		if err == nil {
			_, _ = _this.bodyBuffer.Write(content)
		}
	} else if _this.IsPvBrotli() {
		content, err := _this.BrotliDecompress()
		if err == nil {
			_, _ = _this.bodyBuffer.Write(content)
		}
	}
	if _this.bodyBuffer.Len() == 0 {
		next = false
		err = nil
	} else if err = _this.Decode(_this.bodyBuffer); err != nil {
		_this.bodyBuffer = nil
		next = false
	} else {
		next = true
	}
	return
}
