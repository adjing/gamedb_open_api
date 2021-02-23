package role

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var TableName_CountID = "guid_db"

func APIInitGUIDDB(ctx *gin.Context) {

}

// func APIGETGUIDDB(ctx *gin.Context) {
// 	id := GetIntID("101")
// 	fmt.Println("get int id=", id)

// 	msg := fmt.Sprintf("time: %s id=%d", GetBeiJingTime(), id)
// 	ctx.String(http.StatusOK, msg)
// }

//
func InitSystemIntID(db_name string, guid_key string) int {
	var id = 1
	//
	var dt = Get_Collection(db_name, TableName_CountID)
	var row GUIDB_Com
	row.GUIDKey = guid_key
	row.GUIDInt = 10009

	_, err := dt.InsertOne(context.Background(), row)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func GetInfo_GUIDDB(db_name string, guid_key string) GUIDB_Com {
	var dt = Get_Collection(db_name, TableName_CountID)
	//
	row := GUIDB_Com{}
	cur := dt.FindOne(context.Background(), bson.M{"guid_key": guid_key})
	// cur.Err().Error()
	err2 := cur.Decode(&row)
	fmt.Println(err2)

	return row
}

//guid_int_count.GetIntID
//获取递增ID int
func GetIntID(db_name string, guid_key string) int {
	row1 := GetInfo_GUIDDB(db_name, guid_key)
	if row1.GUIDKey == "" {
		InitSystemIntID(db_name, guid_key)
	}

	fmt.Println(row1)
	//
	var dt = Get_Collection(db_name, TableName_CountID)
	res, err1 := dt.UpdateOne(context.Background(), bson.M{"guid_key": guid_key}, bson.M{"$inc": bson.M{"guid_int": 1}})
	fmt.Println(res, err1)
	//
	row2 := GetInfo_GUIDDB(db_name, guid_key)

	return row2.GUIDInt
}

type GUIDB_Com struct {
	GUIDKey  string `json:"guid_key" bson:"guid_key"`   //1
	GUIDInt  int    `json:"guid_int" bson:"guid_int"`   //2
	FuncText string `json:"func_text" bson:"func_text"` //3
}
