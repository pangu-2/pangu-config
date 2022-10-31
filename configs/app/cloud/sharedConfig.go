package cloud

// "github.com/nacos-group/nacos-sdk-go/vo"

// 服务中心默认 优先拉取
type SharedConfig struct {
	DataId  string `json:"dataId"`
	Group   string `json:"group"`
	Refresh bool   `json:"refresh"`
}

// func SharedConfigToConfigParam(p []SharedConfig) []vo.ConfigParam {
// 	ret := make([]vo.ConfigParam, len(p))
// 	for _, c := range p {
// 		ret = append(ret, vo.ConfigParam{
// 			DataId: c.DataId,
// 			Group:  c.Group,
// 		})
// 	}
// 	return ret
// }
