package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
	"WebSocket/HttpController"
	"WebSocket/Ws"
	"flag"
)



//var ws1 *websocket.Conn
//var ws2 *websocket.Conn
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}
// 路由处理方法 类似各种Controller里的各种Action
func TestHandler(w http.ResponseWriter, r *http.Request ,ps httprouter.Params){
	w.Write([]byte("hhhh"))

	fmt.Println("not found")
}
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))


}


var addr = flag.String("ws", "192.168.31.232:9501", "http service address")
var httpurl = flag.String("http", "192.168.31.232:9501", "http service address")
var start = flag.Int("start", 14, "key start")
var end = flag.Int("end", 15, "key end")
func main() {
	flag.Parse()
	Ws.Init(*addr,*httpurl,*start,*end)
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/CreateBuildReq/:name", HttpController.CreateBuildReq)
	router.POST("/UpdateRoleInfoNameReq/:name",HttpController.UpdateRoleInfoNameReq)
	router.POST("/NpcRelationAdvanceReq/:name",HttpController.NpcRelationAdvanceReq)
	router.POST("/AddNpcFavorabilityReq/:name",HttpController.AddNpcFavorabilityReq)
	router.GET("/NpcFavorabilityReq/:name",HttpController.NpcFavorabilityReq)
	router.POST("/UseItemReq/:name",HttpController.UseItemReq)
	router.POST("/UnlockNpcReq/:name",HttpController.UnlockNpcReq)
	router.POST("/RefStaffReq/:name",HttpController.RefStaffReq)
	router.POST("/ComeOutEmployeeReq/:name",HttpController.ComeOutEmployeeReq)
	router.GET("/LoadStaffReq/:name",HttpController.LoadStaffReq)
	router.POST("/GiveGiftReq/:name",HttpController.GiveGiftReq)
	router.GET("/GiveListReq/:name",HttpController.GiveListReq)
	router.POST("/FriendApplyReq/:name",HttpController.FriendApplyReq)
	router.POST("/FriendAddReq/:name",HttpController.FriendAddReq)

	log.Fatal(http.ListenAndServe(":3001", router))
	//mux := http.NewServeMux()
	//
	//th := &timeHandler{format: time.RFC1123}
	//mux.Handle("/time", th)
	//
	//mux.HandleFunc("/test", 	TestHandler)
	//
	//log.Println("Listening...")
	//http.ListenAndServe(":3000", mux)



}





