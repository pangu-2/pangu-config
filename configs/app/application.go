package app

import "strconv"

//项目信息
type Application struct {
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Version string `json:"version"`
	Debug   bool   `json:"debug"`
}

func (s Application) PortToString() string {
	return strconv.Itoa(s.Port)
}
