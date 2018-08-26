package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

// 刷新商城商品返回 1109
func RefDropShopResult(Uid int32,Data []byte) *AutoMsg.RefDropShopResult{
	RefDropShopResult := &AutoMsg.RefDropShopResult{}
	proto.Unmarshal(Data,RefDropShopResult)
	Param,_ := json.Marshal(RefDropShopResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1109,
		Name:"刷新商城商品返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return RefDropShopResult
}
