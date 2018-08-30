package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type CreateCompany struct {
	Name string
}
//创建公司请求
func CreateCompanyReq(Uid int32,CreateCompany *CreateCompany) []byte{
	CreateCompanyReq := &AutoMsg.CreateCompanyReq{
		Name:CreateCompany.Name,
	}
	Data,_ := proto.Marshal(CreateCompanyReq)
	Param,_ := json.Marshal(CreateCompanyReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1006,
		Name: "创建公司请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1006,Data)
}
