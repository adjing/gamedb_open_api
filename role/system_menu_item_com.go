package role

type SystemMenuItem struct {
	MenuItemGUID   int    `json:"menu_guid" bson:"menu_guid"`               //1 菜单唯一GUID
	MenuItemText   string `json:"menu_text" bson:"menu_text"`               //2 菜单名称
	MenuPath       string `json:"menu_path" bson:"menu_path"`               //3 菜单H5 路径
	MenuIcon       string `json:"menu_icon" bson:"menu_icon"`               //4 菜单图标
	MenuParentGUID int    `json:"menu_parent_guid" bson:"menu_parent_guid"` //5 菜单分类GUID
	MenuParentText string `json:"menu_parent_text" bson:"menu_parent_text"` //6 菜单分类名称
}
