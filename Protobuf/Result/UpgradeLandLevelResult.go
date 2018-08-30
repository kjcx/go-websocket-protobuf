package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//升级土地返回 1171
func UpgradeLandLevelResult(Uid int32,Data []byte) *AutoMsg.UpgradeLandLevelResult {
	UpgradeLandLevelResult:=&AutoMsg.UpgradeLandLevelResult{}
	proto.Unmarshal(Data,UpgradeLandLevelResult)
	Param,_ := json.Marshal(UpgradeLandLevelResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1171,
		Name:  "升级土地返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return UpgradeLandLevelResult
}
