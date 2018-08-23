package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)

//解析商店
func ShopAllResult(Data []byte) *AutoMsg.ShopAllResult {
	ShopAllResult := &AutoMsg.ShopAllResult{}
	err := proto.Unmarshal(Data, ShopAllResult)
	if err != nil {
		fmt.Println("解析ShopAllResult出错")
	}
	return ShopAllResult
}
