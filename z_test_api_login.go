package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_Login(c *gin.Context) {
	var msg = fmt.Sprintf("Game API %s", sys.GetBeiJingTime())
	fmt.Println(msg)
	//

	c.String(http.StatusOK, msg)
}