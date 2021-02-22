package gamedb_open_api

import (
	"encoding/json"
	"fmt"
)

func GetSystemMenuJson(db_name string) {

	lst := GetListAll_SystemMenu(db_name)
	var menutree []SystemMenuTree

	for i := 0; i < len(lst); i++ {
		row := lst[i]
		//
		row2 := SystemMenuTree{}
		row2.MenuGUID = row.MenuGUID
		row2.MenuText = row.MenuText
		row2.MenuIcon = row.MenuIcon

		lstitem := GetListAll_SystemMenuItem(db_name, row.MenuGUID)
		row2.MenuItemList = lstitem
		//
		menutree = append(menutree, row2)
	}
	//
	backcom := GetSendClientComDefault()
	backcom.Data = menutree
	backcom.Text = "suc"
	//
	fmt.Println("menu json:", GetBeiJingTime())
	request_json, err := json.Marshal(backcom)
	if err != nil {

	}
	fmt.Println(string(request_json))
}

func InitSystemData(db_name string) int {

	Insert_LoginAccount_Admin(db_name)

	DeleteAll_SystemMenu(db_name)
	DeleteAll_SystemMenuItem(db_name)
	//
	InitMenu(db_name)
	//
	return 1
}

func GetParentMenuNodeList() []SystemMenu {
	var lst []SystemMenu

	var row_1 = SystemMenu{
		MenuGUID: 101,
		MenuText: "媒体管理",
		MenuPath: "/",
		MenuIcon: "https://gamedb.top/web_img/icon_1.jpg",
	}

	var row_2 = SystemMenu{
		MenuGUID: 102,
		MenuText: "上游管理",
		MenuPath: "/",
		MenuIcon: "https://gamedb.top/web_img/icon_2.jpg",
	}

	var row_3 = SystemMenu{
		MenuGUID: 103,
		MenuText: "数据报表",
		MenuPath: "/",
		MenuIcon: "https://gamedb.top/web_img/icon_3.jpg",
	}

	var row_4 = SystemMenu{
		MenuGUID: 104,
		MenuText: "系统设置",
		MenuPath: "/",
		MenuIcon: "https://gamedb.top/web_img/icon_1.jpg",
	}

	lst = append(lst, row_1)
	lst = append(lst, row_2)
	lst = append(lst, row_3)
	lst = append(lst, row_4)

	return lst
}

func GetParentNode(menu_guid int) SystemMenu {
	var info SystemMenu
	var lst = GetParentMenuNodeList()
	for i := 0; i < len(lst); i++ {
		if lst[i].MenuGUID == menu_guid {
			info = lst[i]
			return info
		}
	}
	return info
}

func GetMenuItemList() []SystemMenuItem {
	var lstItem []SystemMenuItem

	var item_1 = SystemMenuItem{
		MenuItemGUID:   1,
		MenuItemText:   "APP 管理",
		MenuPath:       "/media/app",
		MenuParentGUID: GetParentNode(101).MenuGUID,
		MenuParentText: GetParentNode(101).MenuText,
	}

	var item_2 = SystemMenuItem{
		MenuItemGUID:   2,
		MenuItemText:   "广告位管理",
		MenuPath:       "/ad/page_system_ad_placement",
		MenuParentGUID: GetParentNode(101).MenuGUID,
		MenuParentText: GetParentNode(101).MenuText,
	}

	var item_3 = SystemMenuItem{
		MenuItemGUID:   3,
		MenuItemText:   "广告投放",
		MenuPath:       "/ad/advertising_strategy",
		MenuParentGUID: GetParentNode(101).MenuGUID,
		MenuParentText: GetParentNode(101).MenuText,
	}

	//102-
	var item_4 = SystemMenuItem{
		MenuItemGUID:   4,
		MenuItemText:   "上游信息",
		MenuPath:       "/dsp/index",
		MenuParentGUID: GetParentNode(102).MenuGUID,
		MenuParentText: GetParentNode(102).MenuText,
	}

	var item_5 = SystemMenuItem{
		MenuItemGUID:   5,
		MenuItemText:   "上游广告位",
		MenuPath:       "/dsp/appCode",
		MenuParentGUID: GetParentNode(102).MenuGUID,
		MenuParentText: GetParentNode(102).MenuText,
	}
	//
	//103
	var item_6 = SystemMenuItem{
		MenuItemGUID:   6,
		MenuItemText:   "查询报表",
		MenuPath:       "/media/a",
		MenuParentGUID: GetParentNode(103).MenuGUID,
		MenuParentText: GetParentNode(103).MenuText,
	}

	lstItem = append(lstItem, item_1)
	lstItem = append(lstItem, item_2)
	lstItem = append(lstItem, item_3)
	lstItem = append(lstItem, item_4)
	lstItem = append(lstItem, item_5)
	lstItem = append(lstItem, item_6)

	return lstItem
}

func InitMenu(db_name string) {
	//初始化菜单数据
	fmt.Println("start init data..")

	//
	var lst = GetParentMenuNodeList()
	for i := 0; i < len(lst); i++ {
		Insert_SystemMenu(db_name, lst[i])
	}

	//101-
	var lstItem = GetMenuItemList()
	for j := 0; j < len(lstItem); j++ {
		Insert_SystemMenuItem(db_name, lstItem[j])
	}

	fmt.Println("start init data ok")
}
