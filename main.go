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

func main() {
	flag.Parse()
	Ws.Init()
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/CreateBuildReq/:name", HttpController.CreateBuildReq)
	router.POST("/CreateCompanyReq/:name", HttpController.CreateCompanyReq)
	router.POST("/UpdateRoleInfoNameReq/:name",HttpController.UpdateRoleInfoNameReq)
	router.POST("/NpcRelationAdvanceReq/:name",HttpController.NpcRelationAdvanceReq)
	router.POST("/AddNpcFavorabilityReq/:name",HttpController.AddNpcFavorabilityReq)
	router.GET("/NpcFavorabilityReq/:name",HttpController.NpcFavorabilityReq)
	router.POST("/UseItemReq/:name",HttpController.UseItemReq)
	router.POST("/UnlockNpcReq/:name",HttpController.UnlockNpcReq)//请求解锁npc
	router.POST("/RefStaffReq/:name",HttpController.RefStaffReq)
	router.POST("/ComeOutEmployeeReq/:name",HttpController.ComeOutEmployeeReq)
	router.GET("/LoadStaffReq/:name",HttpController.LoadStaffReq)
	router.POST("/GiveGiftReq/:name",HttpController.GiveGiftReq)
	router.GET("/GiveListReq/:name",HttpController.GiveListReq)
	router.POST("/FriendApplyReq/:name",HttpController.FriendApplyReq)
	router.POST("/FriendAddReq/:name",HttpController.FriendAddReq)
	router.POST("/FriendApplyClearReq/:name",HttpController.FriendApplyClearReq)
	router.POST("/FriendSearchReq/:name",HttpController.FriendSearchReq)
	router.GET("/GetTalentListReq/:name",HttpController.GetTalentListReq)
	router.POST("/BuildLvUpReq/:name",HttpController.BuildLvUpReq)
	router.POST("/AddItem/:name",HttpController.AddItem)
	router.POST("/TalentHireReq/:name",HttpController.TalentHireReq)
	router.GET("/TalentRefreshReq/:name",HttpController.TalentRefreshReq)
	router.POST("/RequestManorReq/:name",HttpController.RequestManorReq)
	router.GET("/ManorVisitInfoReq/:name",HttpController.ManorVisitInfoReq)
	router.POST("/TalentFireReq/:name",HttpController.TalentFireReq)//
	router.POST("/CultivateEmployeeReq/:name",HttpController.CultivateEmployeeReq)//培训员工
	router.POST("/DismissalEmployeeReq/:name",HttpController.DismissalEmployeeReq)//解雇员工
	router.POST("/JoinGameReq/:name",HttpController.JoinGameReq)//加入游戏
	router.POST("/GetMapReq/:name",HttpController.GetMapReq)//大地图
	router.POST("/GetAuctionLandReq/:name",HttpController.GetAuctionLandReq)//获取竞拍土地列表
	router.POST("/AuctionLandReq/:name",HttpController.AuctionLandReq)//竞拍土地请求
	router.POST("/MyLandInfoReq/:name",HttpController.MyLandInfoReq)//我的竞拍土地列表
	router.POST("/FruitsDataReq/:name",HttpController.FruitsDataReq)//水果机列表请求
	router.POST("/RaffleFruitsReq/:name",HttpController.RaffleFruitsReq)//水果机请求
	router.POST("/FriendAddBlackReq/:name",HttpController.FriendAddBlackReq)//拉黑好友
	router.POST("/FriendRemoveBlackReq/:name",HttpController.FriendRemoveBlackReq)//移除黑名单好友
	router.POST("/ReadMailReq/:name",HttpController.ReadMailReq)//阅读邮件请求
	router.POST("/GetMailItemsReq/:name",HttpController.GetMailItemsReq)//一键获取邮件物品请求
	router.POST("/DelMailsReq/:name",HttpController.DelMailsReq)//一键获取邮件物品请求
	router.POST("/RandManorReq/:name",HttpController.RandManorReq)//随机庄园请求
	router.POST("/StealSemenReq/:name",HttpController.StealSemenReq)//偷菜请求
	router.POST("/PublicRoleShopReq/:name",HttpController.PublicRoleShopReq)//获取自己公有区店铺


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





