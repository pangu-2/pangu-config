package pg

// 上传
type Upload struct {
	UploadFolder     string `json:"uploadFolder"`     // 根目录
	StaticAccessPath string `json:"staticAccessPath"` // 相对路径
	SaveLocal        bool   `json:"saveLocal"`        // 保存到本地
	SaveOss          bool   `json:"saveOss"`          // 保存到云存储
}
