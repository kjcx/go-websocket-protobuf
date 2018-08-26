package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//加载npc好感度返回 1037
func NpcFavorabilityResult(Uid int32,Data []byte) *AutoMsg.NpcFavorabilityResult{
	NpcFavorabilityResult := &AutoMsg.NpcFavorabilityResult{}
	proto.Unmarshal(Data,NpcFavorabilityResult)
	Param,_ := json.Marshal(NpcFavorabilityResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1037,
		Name:"加载npc好感度返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return NpcFavorabilityResult
}