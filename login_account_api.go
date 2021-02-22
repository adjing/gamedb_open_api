package gamedb_open_api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func API_GetUser(db_name string, c *gin.Context) {

	lst := GetListAll_LoginAccount(db_name)
	c.JSON(http.StatusOK, lst)
}

func Login(db_name string, r loginRequest) SendClientCom {

	user, err := FindUserByUserName(db_name, r.UserName)
	if err != nil {
		backcom := GetSendClientComDefault()
		backcom.Text = "db error"
		backcom.Code = -1

		return backcom
	}

	if user.UserName == "" {
		backcom := GetSendClientComDefault()
		backcom.Text = "user non-existent"
		backcom.Code = -1

		return backcom
	}

	if user.LoginPassword == "" {
		backcom := GetSendClientComDefault()
		backcom.Text = "密码不能为空"
		backcom.Code = -1

		return backcom
	}

	if user.LoginPassword == r.Passwd {

		token, err := CreateToken(user.UserGUID, JWT_SECRET)
		fmt.Println("token=", token)
		if err != nil {

		}

		backcom := GetSendClientComDefault()
		backcom.Text = "suc"
		backcom.Data = token

		return backcom

	} else {
		backcom := GetSendClientComDefault()
		backcom.Text = "user password error"
		backcom.Code = 4001

		return backcom
	}
}

func Register(db_name string, c *gin.Context) {

	r := registerRequest{}
	err := c.Bind(&r)
	if err != nil {
		return
	}

	if r.Passwd == "" || r.UserName == "" {
		backcom := GetSendClientComDefault()
		backcom.Text = "user name or password cannot empty"
		backcom.Code = 200

		c.JSON(4001, backcom)
		return
	}

	fmt.Println(r)

	var row LoginAccount
	row.UserGUID = GetGUID()
	row.UserName = r.UserName
	row.LoginPassword = r.Passwd
	row.UpdateTime = time.Now().Unix()
	row.UpdateTimeText = GetBeiJingTime()

	err = Insert_LoginAccount(db_name, row)
	if err != nil {
		log.Println(err)
		//
		return
	}
	backcom := GetSendClientComDefault()
	backcom.Text = "suc"

	c.JSON(200, backcom)
}
