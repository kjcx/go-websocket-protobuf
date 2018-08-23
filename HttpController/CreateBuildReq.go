package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"math/rand"
	"fmt"
)

func CreateBuildReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ShopType := rand.Int31n(3) + int32(1)
	Pos := rand.Int31n(3) + int32(1)
	req := Req.CreateBuildReq(&Req.CreateBuild{Pos: Pos, AeraId: 1, ShopType: ShopType})
	fmt.Println(req)
	//go timeWriter(arr[3000],req)
}
