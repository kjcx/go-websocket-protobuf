package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)

type DestoryBuild struct {
	Id string
}
//出售店铺
func DestoryBuildReq(Uid int32,DestoryBuild *DestoryBuild) []byte{
	DestoryBuildReq := &AutoMsg.DestoryBuildReq{
		Id: DestoryBuild.Id,
	}
	Data,_ := proto.Marshal(DestoryBuildReq)
	Param,_ := json.Marshal(DestoryBuildReq)
	log := &Mgo.Log{
		Uid:   Uid,
		MsgId: 1010,
		Name:  "出售店铺请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1010,Data)
}
