package hivesql

type result struct {
	//driver.Result
	affected int64
	lastID   int64
	err      error
}

// RowsAffected returns the number of rows affected by the
// query.
func (r *result) RowsAffected() (int64, error) {
	return r.affected, r.err
}

// LastInsertId returns the database's auto-generated ID
// after, for example, an INSERT into a table with primary
// key.
func (r *result) LastInsertId() (int64, error) {
	return r.lastID, r.err
}
