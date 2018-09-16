package jsonCode

type ParamItem struct {
	Name          string //DTO名称
	Description   string //DTO描述
	Required      string //是否是必须字段
	DTOType       string //DTO类型
	ParamItemList string //DTOType为list时，存储List字段信息
	ListDTOName   string //DTOType为list时，大写DTO文件名称
	PackageName   string //DTOType为list时，Package路径
}
