package gamedb_open_api

type LoginAccount struct {
	UserGUID       string `json:"user_guid" bson:"user_guid"`               //1
	UserID         int    `json:"user_id" bson:"user_id"`                   //2
	UserName       string `json:"user_name" bson:"user_name"`               //3
	LoginPassword  string `json:"login_password" bson:"login_password"`     //4
	StatusCode     string `json:"status_code" bson:"status_code"`           //5
	StatusText     string `json:"status_text" bson:"status_text"`           //6
	UpdateTime     int64  `json:"update_time" bson:"update_time"`           //7
	UpdateTimeText string `json:"update_time_text" bson:"update_time_text"` //8

}

type UserInfo struct {
	UserWXID      string `json:"user_wx_id" `         //1 用户微信ID
	UserName      string `json:"user_name" `          //2 用户名
	LoginPassword int    `json:"mobile_phone_number"` //3 手机号
}

type registerRequest struct {
	UserName string `json:"user_name"`
	Passwd   string `json:"passwd"`
}

type loginRequest struct {
	UserName string `json:"user_name"`
	Passwd   string `json:"passwd"`
}
