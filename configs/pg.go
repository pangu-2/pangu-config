package configs

import "github.com/pangu-2/pangu-config/configs/pg"

type Pg struct {
	Jwt           pg.Jwt           `json:"jwt"`
	Upload        pg.Upload        `json:"upload"`
	Redis         pg.Redis         `json:"redis"`
	Template      pg.Template      `json:"template"`
	Rocketmq      pg.Rocketmq      `json:"rocketmq"`
	Elasticsearch pg.Elasticsearch `json:"elasticsearch"`
}

// 设置默认配置
func (c *Pg) SetDefault() {
	if len(c.Upload.UploadFolder) < 1 {
		c.Upload.UploadFolder = "/uploads"
	}
	if len(c.Upload.UploadFolder) < 1 {
		c.Upload.StaticAccessPath = "/uploads"
	}
}
