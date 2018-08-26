package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)
//提升好感度返回 1133
func AddNpcRelationAdvanceResult(Uid int32,Data []byte) *AutoMsg.AddNpcRelationAdvanceResult{
	AddNpcRelationAdvanceResult:=&AutoMsg.AddNpcRelationAdvanceResult{}
	proto.Unmarshal(Data,AddNpcRelationAdvanceResult)
	fmt.Println(AddNpcRelationAdvanceResult)
	return AddNpcRelationAdvanceResult
}