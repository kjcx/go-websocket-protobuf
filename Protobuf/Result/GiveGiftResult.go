package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)

//赠送礼物返回
func GiveGiftResult(Uid int32,Data []byte) *AutoMsg.GiveGiftResult{
	GiveGiftResult := &AutoMsg.GiveGiftResult{}
	proto.Unmarshal(Data,GiveGiftResult)
	Param,_ := json.Marshal(GiveGiftResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2038,
		Name:  "赠送礼物返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return GiveGiftResult
}