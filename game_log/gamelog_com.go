//1. gamelog_com.go
package game_log

//
type GameLogCom struct {
	Auto_guid     string `json:"auto_guid" bson:"auto_guid"`         //1. pk
	Event_ID      int    `json:"event_id" bson:"event_id"`           //2. 事件ID 1=请求接口 2=响应消息
	Log_time      string `json:"log_time" bson:"log_time"`           //3. 日志时间
	API_Name      string `json:"api_name" bson:"api_name"`           //4. 接口名称
	Send_JSON     string `json:"send_json" bson:"send_json"`         //5. 发送参数
	Response_JSON string `json:"response_json" bson:"response_json"` //6. 响应消息
	Log_type      int    `json:"log_type" bson:"log_type"`           //7. 日志枚举类型 log=1,error=2
	Log_text      string `json:"log_text" bson:"log_text"`           //8. 日志文本内容
	Device_guid   string `json:"device_guid" bson:"device_guid"`     //9. 设备GUID
}
