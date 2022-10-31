package cloud

import (
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// 服务中心配置
type Discovery struct {
	ServerAddr string `json:"serverAddr"`
	Namespace  string `json:"namespace"`
	serverHost string
	serverPort int64
}

func (s Discovery) ServerAddrPort() int64 {
	if len(s.serverHost) < 1 {
		s.ProcessServerAddr()
	}
	return s.serverPort
}

func (s Discovery) ServerAddrHost() string {
	if len(s.serverHost) < 1 {
		s.ProcessServerAddr()
	}
	return s.serverHost
}

// 这里必须使用指针形式,改动值
func (s *Discovery) ProcessServerAddr() {
	host, port, err := net.SplitHostPort(s.ServerAddr)
	log.Infof("host=%v,port=%v", host, port)
	if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
		s.serverPort = 80
	}
	if len(host) > 0 {
		s.serverHost = host
	} else {
		s.serverHost = "localhost"
	}
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		s.serverPort = 80
	} else {
		s.serverPort = p
	}
	log.Infof("host=%v,port=%v", s.serverHost, s.serverPort)
}
