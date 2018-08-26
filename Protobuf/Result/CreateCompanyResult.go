package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

func main() {
	
}
//创建公司结果 1059
func CreateCompanyResult(Uid int32,Data []byte) *AutoMsg.CreateCompanyResult{
	CreateCompanyResult := &AutoMsg.CreateCompanyResult{}
	err := proto.Unmarshal(Data, CreateCompanyResult)
	if err!= nil {
		fmt.Println("解析失败CreateCompanyResult")
	}
	Param,_ := json.Marshal(CreateCompanyResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1059,
		Name:"创建公司结果",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return CreateCompanyResult
}
