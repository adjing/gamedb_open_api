//1. gamelog_com.go
package game_log

//
type GameLogCom struct {
	Auto_guid   string `json:"auto_guid" bson:"auto_guid"`     //1. pk
	Log_type    string `json:"log_type" bson:"log_type"`       //2. 日志枚举类型 log,error
	Log_text    string `json:"log_text" bson:"log_text"`       //3. 日志文本内容
	Log_time    string `json:"log_time" bson:"log_time"`       //4. 日志时间
	Device_guid string `json:"device_guid" bson:"device_guid"` //5. 设备GUID
}
