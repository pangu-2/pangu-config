package pg

import "strconv"

// app.redis
type Redis struct {
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (r Redis) GetAddress() string {
	return r.Host + ":" + r.PortToString()
}

func (r Redis) PortToString() string {
	return strconv.Itoa(r.Port)
}
