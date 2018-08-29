package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type GetMap struct {
	Zone string
	Pos  []int32
}
//1011 请求竞拍土地
func GetMapReq(Uid int32,GetMap *GetMap) []byte{
	GetMapReq := &AutoMsg.GetMapReq{
		Zone: GetMap.Zone,
		Pos:  GetMap.Pos,
	}
	Data,_ := proto.Marshal(GetMapReq)
	Param,_ := json.Marshal(GetMapReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1011,
		Name:  "竞拍土地请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1011,Data)

}
