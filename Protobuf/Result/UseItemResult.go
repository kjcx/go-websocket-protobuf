package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)


//使用道具返回 1078
func UseItemResult(Uid int32,Data []byte) *AutoMsg.UseItemResult{
	UseItemResult := &AutoMsg.UseItemResult{}
	err := proto.Unmarshal(Data, UseItemResult)
	if err != nil {
		log.Fatal("解析数据失败UseItemResult")
	}
	fmt.Println(UseItemResult)
	Param,_ := json.Marshal(UseItemResult)

	log := &Mgo.Log{
		Uid:   Uid,
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
