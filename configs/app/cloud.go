package app

import "github.com/pangu-2/pangu-config/configs/app/cloud"

type Cloud struct {
	Nacos cloud.Nacos `json:"nacos"`
}
