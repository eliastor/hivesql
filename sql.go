package hivesql

import (
	"database/sql"
	"database/sql/driver"
	"net"

	"github.com/beltran/gohive"
)

type drv struct {
}

func (d *drv) Open(dsn string) (driver.Conn, error) {
	cfg, err := parseDSN(dsn)
	if err != nil {
		return nil, err
	}

	ghCfg := gohive.NewConnectConfiguration()

	ghCfg.HiveConfiguration = cfg.hiveConfig
	ghCfg.Username = cfg.user
	ghCfg.Password = cfg.password
	ghCfg.TLSConfig = cfg.tlsConfig
	host, port, _ := net.SplitHostPort(cfg.addr)
	dport, err := net.LookupPort("ip", port)
	if err != nil {
		return nil, err
	}
	ghc, err := gohive.Connect(host, dport, cfg.auth, ghCfg)
	if err != nil {
		return nil, err
	}
	c := new(conn)
	c.c = ghc
	return driver.Conn(c), nil
}

func init() {
	sql.Register("hivesql", &drv{})
}
