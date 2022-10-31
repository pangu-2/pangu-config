package app

import "strconv"

// 数据库信息
type Datasource struct {
	Url      string `json:"url"` //案例 tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
	Username string `json:"username"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (s Datasource) PortToString() string {
	return strconv.Itoa(s.Port)
}
