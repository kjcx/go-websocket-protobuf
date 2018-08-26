package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//提升好感度返回 1133
func AddNpcRelationAdvanceResult(Uid int32,Data []byte) *AutoMsg.AddNpcRelationAdvanceResult{
	AddNpcRelationAdvanceResult:=&AutoMsg.AddNpcRelationAdvanceResult{}
	proto.Unmarshal(Data,AddNpcRelationAdvanceResult)
	fmt.Println(AddNpcRelationAdvanceResult)
	Param,_ := json.Marshal(AddNpcRelationAdvanceResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1133,
		Name:"提升好感度返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return AddNpcRelationAdvanceResult
}