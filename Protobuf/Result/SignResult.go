package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)


//签到列表
func SignResult(Data []byte) *AutoMsg.SignResult  {
	SignResult := &AutoMsg.SignResult{}
	err := proto.Unmarshal(Data, SignResult)
	if err != nil {
		fmt.Println("解析出错")
	}
	return SignResult
}
