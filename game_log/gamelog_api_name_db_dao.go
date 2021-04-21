package game_log

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBCollection_APINameDB = "apiname_db"

//2. MongoDB向集合中插入文档-APINameDB
func Insert_APINameDB(db_name string, row APINameDBCom) error {
	var dt = Get_Collection(db_name, MongoDBCollection_APINameDB)
	_, err := dt.InsertOne(context.Background(), row)
	return err
}

func GetInfo_APINameDB(db_name string, api_name string) APINameDBCom {
	row := APINameDBCom{}

	var dt = Get_Collection(db_name, MongoDBCollection_APINameDB)
	result := dt.FindOne(context.Background(), bson.M{"api_name": api_name})

	result.Decode(&row)

	return row
}

//3. 搜索查询基础方法
func GetListBase_APINameDB(db_name string, filter bson.M) (lst []APINameDBCom, err error) {
	var dt = Get_Collection(db_name, MongoDBCollection_APINameDB)

	var p_pagesize int64 = 100
	var p_pageindex int64 = 1

	opts := new(options.FindOptions)
	limit := p_pagesize
	skip := (p_pageindex - 1) * p_pagesize

	sortMap := make(map[string]interface{})
	sortMap["click_count"] = 1
	opts.Sort = sortMap

	opts.Limit = &limit
	opts.Skip = &skip

	cur, err := dt.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := APINameDBCom{}
		err := cur.Decode(&tmp)
		if err != nil {
			return nil, err
		}
		lst = append(lst, tmp)
	}
	return
}

func GetListAll_APINameDB(db_name string) (lst []APINameDBCom, err error) {
	arr, e := GetListBase_APINameDB(db_name, bson.M{})
	lst = arr
	err = e

	return
}

//4. 根据条件查询列表
// func Get_APINameDB_DayPlacementGUID(ad_placement_guid string, year int, month int, day int) []APINameDBCom {
// 	lst, err := GetListBase_APINameDB(bson.M{"ad_placement_guid": ad_placement_guid, "date_year": year, "date_month": month, "date_day": day})
// 	if err != nil {
// 		//
// 	}
// 	return lst
// }

// //6. 更新
// func Update_APINameDB(doc APINameDBCom) error {
// 	var dt = db.Get_Collection(MongoDBCollection_APINameDB)
// 	_, err := dt.UpdateOne(context.Background(), bson.M{"api_guid": doc.api_guid}, bson.M{"$set": doc})
// 	return err
// }

//7. Delete
func Delete_APINameDB(db_name string, filter bson.M) (count int64, err error) {
	var dt = Get_Collection(db_name, MongoDBCollection_APINameDB)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

//8.
func DeleteAll_APINameDB(db_name string) (count int64, err error) {
	c, e := Delete_APINameDB(db_name, bson.M{})
	return c, e
}

// //9.  分页查询
// func GetListPaging_APINameDB(p_pageindex int64, p_pagesize int64) (lst []APINameDBCom, err error) {

// 	var dt = db.Get_Collection(MongoDBCollection_APINameDB)

// 	if p_pagesize <= 0 {
// 		p_pagesize = 10
// 	}

// 	if p_pageindex <= 0 {
// 		p_pageindex = 1
// 	}

// 	opts := new(options.FindOptions)
// 	limit := p_pagesize
// 	skip := (p_pageindex - 1) * p_pagesize

// 	sortMap := make(map[string]interface{})
// 	sortMap["op_time_text"] = -1
// 	opts.Sort = sortMap

// 	opts.Limit = &limit
// 	opts.Skip = &skip

// 	cur, err := dt.Find(context.Background(), bson.M{}, opts)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	for cur.Next(context.Background()) {
// 		tmp := APINameDBCom{}
// 		err := cur.Decode(&tmp)
// 		if err != nil {
// 			return nil, err
// 		}
// 		lst = append(lst, tmp)
// 	}
// 	return
// }

// click +1
func Update_APINameDB_Click(db_name string, doc APINameDBCom) error {
	var dt = Get_Collection(db_name, MongoDBCollection_APINameDB)
	_, err := dt.UpdateOne(context.Background(), bson.M{"api_guid": doc.Api_guid}, bson.M{"$set": bson.M{"click_count": 1}})
	return err
}

//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo/options"
