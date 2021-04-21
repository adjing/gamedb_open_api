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
	rcom := PagingSelectList{}
	rbyte := GetReceiveComByte(c)

	//2 json to com
	var err = json.Unmarshal(rbyte, &rcom)
	if err != nil {
		fmt.Println(err)
	}

	lst, err := GetListPaging_GameLog(db_name_gamelog, int64(rcom.CurrentPageIndex), 20)
	if err != nil {

	}
	backcom := GetSendClientComDefault()
	backcom.Text = "suc"
	backcom.Data = lst

	c.JSON(http.StatusOK, backcom)
}

func API_getlist_api_name(c *gin.Context) {
	//收到参数
	lst, err := GetListAll_APINameDB(db_name_gamelog)
	if err != nil {

	}
	//
	backcom := GetSendClientComDefault()
	backcom.Text = "suc"
	backcom.Data = lst

	c.JSON(http.StatusOK, backcom)
}

func API_search_by_api_name(c *gin.Context) {
	//收到参数
	rcom := PagingSelectList{}
	rbyte := GetReceiveComByte(c)

	//2 json to com
	var err = json.Unmarshal(rbyte, &rcom)
	if err != nil {
		fmt.Println(err)
	}

	lst, err := GetListPaging_GameLog_Search(db_name_gamelog, 1, 20, rcom.SearchText)
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
	//
	var log = fmt.Sprintf("%s %s", rcom.Log_time, rcom.Log_text)
	fmt.Println(log)

	if rcom.Event_ID == 2 {
		//update api
		dbrow := GetInfo_GameLog_API_Name(db_name_gamelog, rcom.API_Name)
		//
		dbrow.Response_JSON = rcom.Response_JSON
		dbrow.Log_time = sys.GetBeiJingTime()

		Update_APILog(db_name_gamelog, dbrow)

	} else {
		//insert
		err = Insert_GameLog(db_name_gamelog, rcom)
		if err != nil {
			fmt.Println(err)
		}

		//save API
		if rcom.API_Name != "" {
			apiinfo := GetInfo_APINameDB(db_name_gamelog, rcom.API_Name)
			if apiinfo.Api_guid == "" {
				apirow := APINameDBCom{}
				apirow.Api_guid = sys.GetGUID()
				apirow.Api_name = rcom.API_Name
				apirow.Api_desc = apirow.Api_name
				apirow.ClickCount = 1
				//
				Insert_APINameDB(db_name_gamelog, apirow)
			}
		}

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
	// var msg = fmt.Sprintf("del count: %d", count)
	// c.String(http.StatusOK, msg)
	//
	backcom := GetSendClientComDefault()
	backcom.Text = fmt.Sprintf("suc %d", count)
	backcom.Data = nil

	c.JSON(http.StatusOK, backcom)
}

type SendClientCom struct {
	Code int         `json:"status_code"` //1 操作状态ID
	Text string      `json:"status_text"` //2 操作状态描述
	Data interface{} `json:"data"`        //3 返回数据
}

type PagingSelectList struct {
	CurrentPageIndex int    `json:"current_page_index"` //1 当前页索引
	PageSize         int    `json:"page_size"`          //2 页大小 默认20条数据
	SearchText       string `json:"search_text"`        //3 搜索文本
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

	return rev
}
