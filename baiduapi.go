package baiduapi

const (
	BDUAPI_AccessServer_Error = "-50001"
)

type BaiduBiz struct {
	ApiKey string
}

func NewBaiduBiz(apiKey string) BaiduBiz {
	return BaiduBiz{ApiKey: apiKey}
}
