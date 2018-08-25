package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)

//购买时装返回 1023
func ModelClothesResult(Uid int32,Data []byte) *AutoMsg.ModelClothesResult{
	ModelClothesResult := &AutoMsg.ModelClothesResult{}
	proto.Unmarshal(Data,ModelClothesResult)
	Param,_ := json.Marshal(ModelClothesResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1023,
		Name:"购买时装返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)
	return ModelClothesResult
}
