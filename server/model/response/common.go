package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ApiResult struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func InitResult() *ApiResult {
	return &ApiResult{
		Code:  0,
		Msg:   "",
		Data:  nil,
		Count: 0,
	}
}

const (
	Node    = "[]services.Node"
	Cluster = "[]services.Cluster"
	Area    = "[]services.Area"
	Dept    = "[]services.Dept"
	User    = "[]services.User"
)

func (res *ApiResult) SetParams(code int, msg string, data map[string]interface{}) {
	res.Code = code
	res.Msg = msg
	res.Data = data

	if data != nil {
		//		typeStr := reflect.TypeOf(data).String()
		//		log.Printf("uuuu %s\n", typeStr)
		//
		//		if typeStr == Node {
		//			info := data.([]services.Node)
		//		} else if typeStr == Cluster {
		//			info := data.([]services.Cluster)
		//		} else if typeStr == Area {
		//			info := data.([]services.Area)
		//		} else if typeStr == Dept {
		//			info := data.([]services.Dept)
		//		} else if typeStr == User {
		//			info := data.([]services.User)
		//		}
		res.Count = data["num"].(int)
		res.Data = data["list"]
	}
}
