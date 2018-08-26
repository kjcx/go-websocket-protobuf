package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//商店请求
func ShopAllReq(Uid int32,MsgId int32) []byte {
	Shop := &AutoMsg.ShopAllReq{}
	data, err := proto.Marshal(Shop)
	if err != nil {
		fmt.Println("1")
	}
	Param,_ := json.Marshal(Shop)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: MsgId,
		Name: "商店请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(MsgId, data)

}
