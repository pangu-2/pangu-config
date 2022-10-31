package configs

import app2 "github.com/pangu-2/pangu-config/configs/app"

type App struct {
	Application app2.Application `json:"application"`
	Datasource  app2.Datasource  `json:"datasource"`
	Query       app2.Query       `json:"query"`
	Cloud       app2.Cloud       `json:"cloud"`
}

// 设置默认配置
func (c *App) SetDefault() {
	//查询缓存 存储
	c.Query.ResultCache.CacheStore = "REDIS"
}
