package gamedb_open_api

//2020-12-15
import (
	"context"

	// "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

var TableName_LoginAccount = "login_account"

//database dao
func Insert_LoginAccount_Admin(db_name string) {

	var UserName = "admin"
	//
	Delete_LoginAccount_UserName(db_name, UserName)
	//
	var row LoginAccount

	row.UserGUID = GetGUID()
	row.UserName = UserName
	row.LoginPassword = "123456"
	row.UpdateTime = GetNowTimestamp()
	row.UpdateTimeText = GetBeiJingTime()
	//
	Insert_LoginAccount(db_name, row)
}

func Insert_LoginAccount(db_name string, row LoginAccount) error {

	var dt = Get_Collection(db_name, TableName_LoginAccount)
	_, err := dt.InsertOne(context.Background(), row)
	return err
}

func FindUserByUserName(db_name string, user_name string) (row LoginAccount, err error) {
	lst, err := GetList_LoginAccount(db_name, bson.M{"user_name": user_name})

	if err != nil {

	}

	if len(lst) > 0 {
		row = lst[0]
	}

	return row, err
}

func GetListAll_LoginAccount(db_name string) (pro []LoginAccount) {
	lst, err := GetList_LoginAccount(db_name, bson.M{})
	if err != nil {

	}

	return lst

}

func GetList_LoginAccount(db_name string, filter bson.M) (pro []LoginAccount, err error) {
	var dt = Get_Collection(db_name, TableName_LoginAccount)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		//log
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := LoginAccount{}
		err := cur.Decode(&tmp)
		if err != nil {
			//log
			return nil, err
		}
		pro = append(pro, tmp)
	}
	return
}

//7. Delete
func Delete_LoginAccount(db_name string, filter bson.M) (count int64, err error) {
	var dt = Get_Collection(db_name, TableName_LoginAccount)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

//8.
func Delete_LoginAccount_UserName(db_name string, user_name string) (count int64, err error) {
	c, e := Delete_LoginAccount(db_name, bson.M{"user_name": user_name})
	return c, e
}
