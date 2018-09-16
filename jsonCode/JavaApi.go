package jsonCode

type JavaApi struct {
	UrlPath            string
	ControllerName     string
	MethodName         string
	HTTPMethod         string //HTTP请求方法
	RequestDTO         []ParamItem
	ResponseDTO        []ParamItem
	PackageName        string
	RequestImportData  []string //RequestDTO需要Import的类型,比如Date, List
	ResponseImportData []string //ResponseDTO需要Import的类型,比如Date, List
}
