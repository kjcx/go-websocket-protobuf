package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//自己公有区店铺返回 1192
func PublicRoleShopResult(Uid int32,Data []byte) *AutoMsg.PublicRoleShopResult  {
	PublicRoleShopResult := &AutoMsg.PublicRoleShopResult{}
	proto.Unmarshal(Data,PublicRoleShopResult)
	Param,_ := json.Marshal(PublicRoleShopResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1192,
		Name:  "自己公有区店铺返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return  PublicRoleShopResult
}
