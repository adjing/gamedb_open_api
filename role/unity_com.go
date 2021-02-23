package role

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//s-c返回给客户端的统一组件
type SendClientCom struct {
	Code int         `json:"status_code"` //1 操作状态ID
	Text string      `json:"status_text"` //2 操作状态描述
	Data interface{} `json:"data"`        //3 返回数据
}

//client - server
type UnityClientCom struct {
	UserGUID string `json:"user_guid"` //1 用户UserGUID
	CardGUID string `json:"card_guid"` //2 卡片编号
	// UserGUID_1 string `json:"user_guid_1"` //3 用户UserGUID 1
	// UserGUID_2 string `json:"user_guid_2"` //4 用户UserGUID 2
	AutoGUID_1 string `json:"player_card_auto_guid_1"` //5 卡片编号 1
	AutoGUID_2 string `json:"player_card_auto_guid_2"` //6 卡片编号 2
}

func GetSendClientComDefault() SendClientCom {

	com := SendClientCom{}
	com.Code = http.StatusOK

	return com
}

//收到组件 SystemAdPlacement
func GetReceiveComByte(c *gin.Context) []byte {
	//收到参数
	rev, err := ioutil.ReadAll(c.Request.Body)
	//
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("rev_com text =")

	txt := string(rev)
	fmt.Println(txt)
	// fmt.Println(rev_com)

	return rev
}
