package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"encoding/json"
	"time"
)
//出售道具返回 1069
func SellItemResult(Uid int32 ,Data []byte) *AutoMsg.SellItemResult {
	SellItemResult := &AutoMsg.SellItemResult{}
	proto.Unmarshal(Data,SellItemResult)
	Param,_ := json.Marshal(SellItemResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1069,
		Name:"出售道具返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)
	return SellItemResult
}
