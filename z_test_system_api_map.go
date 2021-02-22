package gamedb_open_api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func InitGinRoute() {

	route = gin.Default()

	//
	route.GET("/", HomePage)
	v1 := route.Group("v1")

	v1.POST("/login", API_Login)

	var port = 9095
	fmt.Println("run 9095 time:", GetBeiJingTime())
	route.Run(fmt.Sprintf(":%d", port))

}

func HomePage(c *gin.Context) {
	var msg = fmt.Sprintf("Game API %s", GetBeiJingTime())
	fmt.Println(msg)
	//

	c.String(http.StatusOK, msg)
}
