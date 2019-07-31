package hivesql

import (
	"context"
	"database/sql/driver"
	"strings"

	"github.com/beltran/gohive"
)

type conn struct {
	c *gohive.Connection
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	q := strings.TrimRight(query, ";")
	st := newStmt(q, c)
	return driver.Stmt(st), nil
}

func (c *conn) Close() error {
	c.c.Close()
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	panic("not implemented")
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	panic("not implemented")
}

// // Implemetation of Queryer interface
// func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
// 	//remove last semicolon
// 	q := strings.TrimRight(query, ";")

// 	st := newStmt(q)
// 	return st.QueryContext(ctx, args)

// }
