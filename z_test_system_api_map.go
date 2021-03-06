package main

import (
	"fmt"
	"gamedb_open_api/game_log"
	"gamedb_open_api/role"
	"gamedb_open_api/sys"

	"net/http"

	// "github.com/adjing/gamedb_open_api/role"
	// "github.com/adjing/gamedb_open_api/sys"
	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func InitGinRoute() {

	route = gin.Default()

	//
	route.GET("/", HomePage)
	v1 := route.Group("v1")

	// v1.POST("/login", API_Login)

	//role.InitSystemData
	v1.GET("/initsys", API_InitSystemData)    //test
	v1.GET("/getmenu", API_GetSystemMenuJson) //test
	v1.POST("/insertone_gamelog", game_log.APIInsertOne_GameLog)
	v1.POST("/deleteall_gamelog", game_log.APIDeleteAll_GameLog)
	v1.GET("/deleteall_gamelog", game_log.APIDeleteAll_GameLog)

	v1.POST("/getlist_gamelog", game_log.APIGetList_GameLog)
	v1.POST("/getlist_gamelog_error", game_log.APIGetList_GameLog_Error)

	v1.POST("/search_by_api_name", game_log.API_search_by_api_name)
	v1.POST("/getlist_api_name", game_log.API_getlist_api_name)

	var port = 9095
	fmt.Println("run 9095 time:", sys.GetBeiJingTime())
	route.Run(fmt.Sprintf(":%d", port))

}

func API_GetSystemMenuJson(c *gin.Context) {
	var b = role.GetSystemMenuJson(m_db_name)

	c.JSON(http.StatusOK, b)
}

func API_InitSystemData(c *gin.Context) {
	role.InitSystemData(m_db_name)

	c.String(http.StatusOK, "ok")
}

func HomePage(c *gin.Context) {
	var msg = fmt.Sprintf("Game API %s", sys.GetBeiJingTime())
	fmt.Println(msg)
	//

	c.String(http.StatusOK, msg)
}
