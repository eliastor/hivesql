package hivesql

import (
	"database/sql/driver"
)

type conn struct {
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {

	return nil, nil
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return driver.Tx(new(tx)), nil
}
