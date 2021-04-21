package game_log

type APINameDBCom struct {
	Api_guid   string `json:"api_guid" bson:"api_guid"`       //1. api_guid
	Api_name   string `json:"api_name" bson:"api_name"`       //2. api_name
	Api_desc   string `json:"api_desc" bson:"api_desc"`       //3. api_desc
	ClickCount int    `json:"click_count" bson:"click_count"` //3. count
}
