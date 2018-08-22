package Result

import (
	"github.com/golang/protobuf/proto/test"
	"github.com/golang/protobuf/proto"
	"log"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)


//使用道具返回 1078
func UseItemResult(Data []byte) *AutoMsg.UseItemResult{
	UseItemResult := &AutoMsg.UseItemResult{}
	err := proto.Unmarshal(Data, UseItemResult)
	if err != nil {
		log.Fatal("解析数据失败UseItemResult")
	}
	Param,_ := json.Marshal(UseItemResult)

	log := &Mgo.Log{
		Uid:   14,
		MsgId: 1078,
		Name:  "使用道具返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)
	return UseItemResult
}
