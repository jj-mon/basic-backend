package config

import "net"

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := "localhost"

	port := "8000"

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
