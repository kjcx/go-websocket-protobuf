package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//2026解锁npc返回
func UnlockNpcResult(Uid int32,Data []byte) *AutoMsg.UnlockNpcResult{
	UnlockNpcResult := &AutoMsg.UnlockNpcResult{}
	proto.Unmarshal(Data,UnlockNpcResult)
	Param,_ := json.Marshal(UnlockNpcResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:2026,
		Name:"解锁npc返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return UnlockNpcResult
}
