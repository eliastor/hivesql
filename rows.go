package hivesql

import (
	"context"
	"database/sql/driver"
	"io"

	"github.com/beltran/gohive"
)

type rows struct {
	schema  map[string]string
	cursor  *gohive.Cursor
	columns []string
	ctx     context.Context
}

func newRow(ctx context.Context, cursor *gohive.Cursor) *rows {
	r := new(rows)
	r.ctx = ctx
	r.cursor = cursor
	return r
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *rows) Columns() []string {
	if r.schema == nil {
		r.schema = make(map[string]string)

		//r.cursor.Poll(false)

		for _, col := range r.cursor.Description() {
			//[]string{column.ColumnName, typeDesc.PrimitiveEntry.Type.String()}
			r.schema[col[0]] = col[1]
			r.columns = append(r.columns, col[0])
		}
	}
	return r.columns
}

// Close closes the rows iterator.
func (r *rows) Close() error {
	r.cursor.Close()
	return nil
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (r *rows) Next(dest []driver.Value) error {
	if r.cursor.HasMore(r.ctx) {
		var err error
		if err = r.cursor.Error(); err != nil {
			return err
		}
		rMap := r.cursor.RowMap(r.ctx)
		if err = r.cursor.Error(); err != nil {
			return err
		}
		for k, name := range r.Columns() {
			dest[k] = rMap[name]
		}
		return nil
	}
	return io.EOF
}
