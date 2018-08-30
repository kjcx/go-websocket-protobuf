package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//1118 加载员工返回
func LoadStaffResult(Uid int32,Data []byte) *AutoMsg.LoadStaffResult {
	LoadStaffResult := &AutoMsg.LoadStaffResult{}
	proto.Unmarshal(Data,LoadStaffResult)
	Param,_ := json.Marshal(LoadStaffResult)
	fmt.Println(Param)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1118,
		Name:"加载员工数据返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return LoadStaffResult
}
