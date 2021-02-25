package game_log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gamedb_open_api/sys"

	"github.com/gin-gonic/gin"
)

//gamelog_api.go
var db_name_gamelog = "game_dev_log"

//v1.POST("/getlist_gamelog", game_log.APIGetList_GameLog)
func APIGetList_GameLog(c *gin.Context) {
	//收到参数

	lst, err := GetListPaging_GameLog(db_name_gamelog, 1, 100)
	if err != nil {

	}
	backcom := GetSendClientComDefault()
	backcom.Text = "suc"
	backcom.Data = lst

	c.JSON(http.StatusOK, backcom)
}

//v1.POST("/insertone_gamelog", game_log.APIInsertOne_GameLog)
func APIInsertOne_GameLog(c *gin.Context) {
	//1 Server and client communication components
	rcom := GameLogCom{}
	rbyte := GetReceiveComByte(c)

	//2 json to com
	var err = json.Unmarshal(rbyte, &rcom)
	if err != nil {
		fmt.Println(err)
	}

	//3 Is it in the database

	//4 set value
	rcom.Auto_guid = sys.GetGUID()
	rcom.Log_time = sys.GetBeiJingTime()

	//insert
	err = Insert_GameLog(db_name_gamelog, rcom)
	if err != nil {
		fmt.Println(err)
	}
	backcom := GetSendClientComDefault()
	backcom.Text = "suc"
	backcom.Data = fmt.Sprintf("suc %s", sys.GetBeiJingTime())
	c.JSON(http.StatusOK, backcom)
}

//v1.POST("/deletestatus_gamelog", game_log.APIDeleteStatus_GameLog)
func APIDeleteAll_GameLog(c *gin.Context) {
	//1 Server and client communication components
	count, err := DeleteAll_GameLog(db_name_gamelog)
	if err != nil {

	}
	var msg = fmt.Sprintf("del count: %d", count)
	c.String(http.StatusOK, msg)
}

type SendClientCom struct {
	Code int         `json:"status_code"` //1 操作状态ID
	Text string      `json:"status_text"` //2 操作状态描述
	Data interface{} `json:"data"`        //3 返回数据
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
