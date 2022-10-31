package pg

import l "github.com/pangu-2/pangu-config/configs/pg/rocketmq"

// 消息队列
type Rocketmq struct {
	NameServer  string     `json:"name-server"`
	NameServers []string   `json:"nameServers"`
	Producer    l.Producer `json:"producer"`
}

// 设置默认配置
func (c *Rocketmq) SetDefault() {
	//	if len(c.NameServer) < 1 {
	//		c.NameServer = "localhost:9876"
	//	}
	//
	//	if len(c.NameServers) < 1 {
	//		c.NameServers = append(c.NameServers, c.NameServer)
	//	}
	//
	//	if len(c.Producer.Group) < 1 {
	//		c.Producer.Group = "default"
	//	}
}
