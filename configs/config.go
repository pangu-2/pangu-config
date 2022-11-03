package configs

import (
	app2 "github.com/pangu-2/pangu-config/configs/app"
	"github.com/pangu-2/pangu-config/configs/app/cloud"
	"github.com/pangu-2/pangu-config/configs/pg"
)

type ConfigPg struct {
	App App
	Pg  Pg
}

var conf map[string]interface{}
var configPg ConfigPg

// 需要获取配置列表
var confRegistryList []cloud.SharedConfig

// 获取配置后保存的原始配置数据
var confRegistrySource map[string]string

// 设置默认配置
func (c *ConfigPg) SetDefault() {
	//模版引擎 PONGO2,TEMPLATE(TEMPLATE Default)
	//模版引擎 PONGO2,TEMPLATE(TEMPLATE Default)
	c.Pg.Template.Engine = "PONGO2"
	c.Pg.Template.Type = "PONGO2"
	c.Pg.Template.FileData = "FILE"
	c.Pg.Template.Path = "template/pongo2"
	c.Pg.Template.Suffix = ".html"
}

func GetConfig() App {
	return configPg.App
}
func GetConfigMap() map[string]interface{} {
	return conf
}

func GetApplication() app2.Application {
	return configPg.App.Application
}

func GetDatasource() app2.Datasource {
	return configPg.App.Datasource
}

func GetCloudDiscovery() cloud.Discovery {
	return configPg.App.Cloud.Nacos.Discovery
}

func GetQuery() app2.Query {
	return configPg.App.Query
}
func GetRedis() pg.Redis {
	return configPg.Pg.Redis
}

func GetTemplate() pg.Template {
	return configPg.Pg.Template
}

func GetPg() Pg {
	return configPg.Pg
}

func GetPgUpload() pg.Upload {
	return configPg.Pg.Upload
}

func GetPgJwt() pg.Jwt {
	return configPg.Pg.Jwt
}

func GetRocketmq() pg.Rocketmq {
	return configPg.Pg.Rocketmq
}

func GetEs() pg.Elasticsearch {
	return configPg.Pg.Elasticsearch
}
