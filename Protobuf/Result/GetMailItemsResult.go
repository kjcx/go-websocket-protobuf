package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//一键获取邮件物品返回 1095
func GetMailItemsResult(Uid int32,Data []byte) *AutoMsg.GetMailItemsResult {
	GetMailItemsResult := &AutoMsg.GetMailItemsResult{}
	proto.Unmarshal(Data,GetMailItemsResult)
	Param,_ := json.Marshal(GetMailItemsResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1095,
		Name:  "一键获取邮件物品返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return GetMailItemsResult
}
