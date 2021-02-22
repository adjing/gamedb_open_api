package gamedb_open_api

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var TableName_SystemMenuItem = "system_menu_item"

//创建一条新记录
func Insert_SystemMenuItem(db_name string, row SystemMenuItem) error {

	var dt = Get_Collection(db_name, TableName_SystemMenuItem)
	_, err := dt.InsertOne(context.Background(), row)
	return err
}

func GetList_SystemMenuItem(db_name string, filter bson.M) (pro []SystemMenuItem, err error) {
	var dt = Get_Collection(db_name, TableName_SystemMenuItem)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		//log
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := SystemMenuItem{}
		err := cur.Decode(&tmp)
		if err != nil {
			//log
			return nil, err
		}
		pro = append(pro, tmp)
	}
	return
}

func GetSystemMenuItemInfo(menu_guid int) []SystemMenuItem {

	var lst []SystemMenuItem
	return lst
}

func GetListAll_SystemMenuItem(db_name string, menu_guid int) []SystemMenuItem {
	lst, err := GetList_SystemMenuItem(db_name, bson.M{"menu_parent_guid": menu_guid})
	if err != nil {
	}

	return lst
}

//7. Delete
func Delete_SystemMenuItem(db_name string, filter bson.M) (count int64, err error) {
	var dt = Get_Collection(db_name, TableName_SystemMenuItem)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

//8.
func DeleteAll_SystemMenuItem(db_name string) (count int64, err error) {
	c, e := Delete_SystemMenuItem(db_name, bson.M{})
	return c, e
}
