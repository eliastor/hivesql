package hivesql

import (
	"crypto/tls"
	"crypto/x509"
	"net"
)

type config struct {
	addr     string
	user     string
	password string
	dbName   string

	auth string

	hiveConfig map[string]string
	tlsConfig  *tls.Config
}

func (c *config) turnTLS() {
	if c.tlsConfig == nil {
		c.tlsConfig = new(tls.Config)
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			rootCAs = x509.NewCertPool()
		}
		c.tlsConfig.RootCAs = rootCAs
	}
}

func (c *config) normalize() {
	if c.addr == "" {
		c.addr = "localhost:10000"
	}
	c.addr = ensureHavePort(c.addr)
	if c.dbName == "" {
		c.dbName = "default"
	}
	if c.auth == "" {
		c.auth = "NONE"
	}
}

func ensureHavePort(addr string) string {
	if _, _, err := net.SplitHostPort(addr); err != nil {
		return net.JoinHostPort(addr, "10000")
	}
	return addr
}
