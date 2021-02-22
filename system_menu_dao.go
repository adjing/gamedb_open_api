package gamedb_open_api

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var TableName_SystemMenu = "system_menu"

//创建一条新记录 database_mongodb.Insert_SystemMenu
func Insert_SystemMenu(db_name string, row SystemMenu) error {

	var dt = Get_Collection(db_name, TableName_SystemMenu)
	_, err := dt.InsertOne(context.Background(), row)
	return err
}

func GetList_SystemMenu(db_name string, filter bson.M) (pro []SystemMenu, err error) {
	var dt = Get_Collection(db_name, TableName_SystemMenu)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		//log
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := SystemMenu{}
		err := cur.Decode(&tmp)
		if err != nil {
			//log
			return nil, err
		}
		pro = append(pro, tmp)
	}
	return
}

func GetListAll_SystemMenu(db_name string) []SystemMenu {
	lst, err := GetList_SystemMenu(db_name, bson.M{})
	if err != nil {
	}

	return lst
}

//7. Delete
func Delete_SystemMenu(db_name string, filter bson.M) (count int64, err error) {
	var dt = Get_Collection(db_name, TableName_SystemMenu)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

//8.
func DeleteAll_SystemMenu(db_name string) (count int64, err error) {
	c, e := Delete_SystemMenu(db_name, bson.M{})
	return c, e
}
