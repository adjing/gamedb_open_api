//1. gamelog_dao.go
package game_log

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBCollection_GameLog = "game_log"

//2. MongoDB向集合中插入文档-GameLog
func Insert_GameLog(db_name string, row GameLogCom) error {
	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)
	_, err := dt.InsertOne(context.Background(), row)
	return err
}

//3. 搜索查询基础方法
func GetListBase_GameLog(db_name string, filter bson.M) (pro []GameLogCom, err error) {
	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		//log
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := GameLogCom{}
		err := cur.Decode(&tmp)
		if err != nil {
			//log
			return nil, err
		}
		pro = append(pro, tmp)
	}
	return
}

// //4. 根据条件查询列表
// func Get_GameLog_DayPlacementGUID(ad_placement_guid string, year int, month int, day int) []GameLogCom {
// 	lst, err := GetListBase_GameLog(bson.M{"ad_placement_guid": ad_placement_guid, "date_year": year, "date_month": month, "date_day": day})
// 	if err != nil {
// 		//
// 	}
// 	return lst
// }

//5. 查询一条记录(一个document)
func GetInfo_GameLog_API_Name(db_name string, api_name string) GameLogCom {
	lst, err := GetListBase_GameLog(db_name, bson.M{"api_name": api_name})
	var count = len(lst)
	if err != nil {
		//
	}
	var row GameLogCom
	if count > 0 {
		row = lst[0]
	}
	return row
}

//6. 更新
func Update_APILog(db_name string, doc GameLogCom) error {
	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)
	_, err := dt.UpdateOne(context.Background(), bson.M{"api_name": doc.API_Name}, bson.M{"$set": doc})
	return err
}

//7. Delete
func Delete_GameLog(db_name string, filter bson.M) (count int64, err error) {
	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

//8.
func DeleteAll_GameLog(db_name string) (count int64, err error) {
	c, e := Delete_GameLog(db_name, bson.M{})
	return c, e
}

//9.  分页查询
func GetListPaging_GameLog(db_name string, p_pageindex int64, p_pagesize int64) (lst []GameLogCom, err error) {

	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)

	if p_pagesize <= 0 {
		p_pagesize = 10
	}

	if p_pageindex <= 0 {
		p_pageindex = 1
	}

	opts := new(options.FindOptions)
	limit := p_pagesize
	skip := (p_pageindex - 1) * p_pagesize

	sortMap := make(map[string]interface{})
	sortMap["log_time"] = -1
	opts.Sort = sortMap

	opts.Limit = &limit
	opts.Skip = &skip

	cur, err := dt.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := GameLogCom{}
		err := cur.Decode(&tmp)
		if err != nil {
			return nil, err
		}
		lst = append(lst, tmp)
	}
	return
}

//9.  分页查询 Search
func GetListPaging_GameLog_Search(db_name string, p_pageindex int64, p_pagesize int64, api_name string) (lst []GameLogCom, err error) {

	var dt = Get_Collection(db_name, MongoDBCollection_GameLog)

	if p_pagesize <= 0 {
		p_pagesize = 10
	}

	if p_pageindex <= 0 {
		p_pageindex = 1
	}

	opts := new(options.FindOptions)
	limit := p_pagesize
	skip := (p_pageindex - 1) * p_pagesize

	sortMap := make(map[string]interface{})
	sortMap["log_time"] = -1
	opts.Sort = sortMap

	opts.Limit = &limit
	opts.Skip = &skip

	cur, err := dt.Find(context.Background(), bson.M{"api_name": api_name}, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := GameLogCom{}
		err := cur.Decode(&tmp)
		if err != nil {
			return nil, err
		}
		lst = append(lst, tmp)
	}
	return
}

func GetListAll_GameLogCom(db_name string) (lst []GameLogCom, err error) {

	lst1, err1 := GetListBase_GameLog(db_name, bson.M{})
	var count = len(lst)
	if err1 != nil {
		//
		err = err1
	}

	if count > 0 {
		lst = lst1
	}

	return

}

//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo/options"
