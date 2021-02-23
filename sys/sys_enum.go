package sys

//import (
//	"microsecond_tech_ads_2020/microsecond_tech_ads_sys/db"
//	"context"
//)

/*
import (
	"microsecond_tech_ads_2020/microsecond_tech_ads_sys/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"adstore/slog"
)

var DataTableName_Enum_System = "enum_system"


/////////////////////dao
func GetEnumSystem(filter bson.M) (pro []Enum_SystemInfo, err error) {

	// c := db.GetClient()
	var dt = db.Get_Collection(DataTableName_Enum_System)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		slog.Logger.Error(err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := Enum_SystemInfo{}
		err := cur.Decode(&tmp)
		if err != nil {
			slog.Logger.Error(err)
			return nil, err
		}
		pro = append(pro, tmp)
	}

	return
}

func Insert_EnumSystem(in []Enum_SystemInfo) error {
	var in_interface []interface{}

	for k, _ := range in {
		in_interface = append(in_interface, in[k])
	}
	//
	var dt = db.Get_Collection(DataTableName_Enum_System)
	_, err := dt.InsertMany(context.Background(), in_interface)

	return err

}

func GetInfo_EnumSystem_EnumGUID(enumguid string) Enum_SystemInfo {

	lst, err := GetEnumSystem(bson.M{"enumguid": enumguid})
	var count = len(lst)
	if err != nil {

	}
	//
	var row Enum_SystemInfo
	if count > 0 {
		row = lst[0]
	}
	return row
}

func InitData_EnumSystem() []Enum_SystemInfo {
	var lst []Enum_SystemInfo

	var row Enum_SystemInfo
	row.EnumGUID = "1"
	row.Summary = "设备类型"

	//
	var enumItemList []EnumItemInfo

	var item_1 EnumItemInfo
	item_1.ItemGUID = GetGUID()
	item_1.ItemCodeName = "Android"
	item_1.ItemSummary = "Android"
	item_1.ItemIntValue = 1
	//
	var item_2 EnumItemInfo
	item_2.ItemGUID = GetGUID()
	item_2.ItemCodeName = "iOS"
	item_2.ItemSummary = "iOS"
	item_2.ItemIntValue = 2
	//
	var item_3 EnumItemInfo
	item_3.ItemGUID = GetGUID()
	item_3.ItemCodeName = "Other"
	item_3.ItemSummary = "其他"
	item_3.ItemIntValue = 3

	enumItemList = append(enumItemList, item_1)
	enumItemList = append(enumItemList, item_2)
	enumItemList = append(enumItemList, item_3)

	//
	row.EnumItemList = enumItemList

	lst = append(lst, row)
	//
	Insert_EnumSystem(lst)

	return lst
}

*/
