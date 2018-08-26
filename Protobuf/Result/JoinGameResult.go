package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

//加入游戏返回 1066
func JoinGameResult(Data []byte) *AutoMsg.JoinGameResult {
	JoinGameResult := &AutoMsg.JoinGameResult{}
	err := proto.Unmarshal(Data, JoinGameResult)
	if err != nil {
		fmt.Println("marshaling error:1012 ", err.Error())
	}

	return JoinGameResult
}
