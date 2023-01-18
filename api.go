package bili_live_ws_codec

import "encoding/json"

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
	_this.ProtocolVersion = PvPopular
	_this.Operation = OpJoinRoom
	_this.Body = body
	return
}

func (_this *Packet) JoinRoom2(roomId json.Number, key string) (err error) {
	return _this.JoinRoom("0", roomId, 3, "web", 2, key)
}
