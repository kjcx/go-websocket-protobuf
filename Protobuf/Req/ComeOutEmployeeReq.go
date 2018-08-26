package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type ComeOutEmployee struct {
	ShopId               string // 店铺id
	ComeOutInEmployee    int32 //
	NpcCardId            []string //员工id
}
//调入调出员工 1091 调入(0)还是调出(1)
func ComeOutEmployeeReq(Uid int32,ComeOutEmployee *ComeOutEmployee) []byte {
	ComeOutEmployeeReq := &AutoMsg.ComeOutEmployeeReq{
		ShopId:            ComeOutEmployee.ShopId,
		ComeOutInEmployee: ComeOutEmployee.ComeOutInEmployee,
		NpcCardId:         ComeOutEmployee.NpcCardId,
	}
	Data,_ := proto.Marshal(ComeOutEmployeeReq)
	Param,_ := json.Marshal(LoadStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1091,
		Name: "调入调出员工请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1091,Data)
}
