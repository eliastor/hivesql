package hivesql

import (
	"crypto/tls"
	"net"
)

type config struct {
	addr       string
	auth       string
	user       string
	password   string
	dbName     string
	hiveConfig map[string]string
	tlsConfig  *tls.Config
}

func (c *config) normalize() {
	if c.addr == "" {
		c.addr = "localhost:10000"
	}
	c.addr = ensureHavePort(c.addr)
	if c.dbName == "" {
		c.dbName = "default"
	}
}

func ensureHavePort(addr string) string {
	if _, _, err := net.SplitHostPort(addr); err != nil {
		return net.JoinHostPort(addr, "10000")
	}
	return addr
}
