package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//创建店铺结构
type CreateBuild struct {
	Pos int32
	AeraId int32
	ShopType int32
}
//发送创建店铺 1005
func CreateBuildReq(Uid int32,CreateBuild *CreateBuild)[]byte {
	CreateBuildReq := &AutoMsg.CreateBuildReq{
		Pos:CreateBuild.Pos,
		AreaId:CreateBuild.AeraId,
		ShopType:CreateBuild.ShopType,
	}
	Data,_ := proto.Marshal(CreateBuildReq)
	Param,_ := json.Marshal(CreateBuildReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1005,
		Name: "创建店铺请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1005,Data)

}
