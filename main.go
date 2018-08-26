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
func TestHandler(w http.ResponseWriter, r *http.Request ,){
	w.Write([]byte("hhhh"))

	fmt.Println("not found")
}
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))


}


var addr = flag.String("ws", "192.168.31.64:9501", "http service address")
var httpurl = flag.String("http", "192.168.31.64:9501", "http service address")
var start = flag.Int("start", 14, "key start")
var end = flag.Int("end", 15, "key end")
func main() {
	flag.Parse()
	Ws.Init(*addr,*httpurl,*start,*end)
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/CreateBuildReq/:Pos/:AreaId/:ShopType", HttpController.CreateBuildReq)
	router.GET("/UpdateRoleInfoNameReq/:name",HttpController.UpdateRoleInfoNameReq)

	log.Fatal(http.ListenAndServe(":3000", router))
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





