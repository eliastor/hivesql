package hivesql

import (
	"database/sql/driver"
)

type drv struct {
}

func (d *drv) Open(name string) (driver.Conn, error) {
	return driver.Conn(new(conn)), nil
}
