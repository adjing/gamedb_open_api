package gamedb_open_api

type SystemMenu struct {
	MenuGUID int    `json:"menu_guid" bson:"menu_guid"` //1. 菜单分类GUID
	MenuText string `json:"menu_text" bson:"menu_text"` //2. 菜单分类名称
	MenuPath string `json:"menu_path" bson:"menu_path"` //3 菜单H5 路径
	MenuIcon string `json:"menu_icon" bson:"menu_icon"` //4 菜单图标

}

func (u *SystemMenu) TableName() string {
	return "system_menu"
}

type SystemMenuTree struct {
	MenuGUID     int              `json:"menu_guid"`      //1. 菜单分类GUID
	MenuText     string           `json:"menu_text"`      //2. 菜单分类名称
	MenuPath     string           `json:"menu_path"`      //3 菜单H5 路径
	MenuIcon     string           `json:"menu_icon"`      //4 菜单图标
	MenuItemList []SystemMenuItem `json:"menu_item_list"` //4. 子菜单列表
}
