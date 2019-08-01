package hivesql

import (
	"context"
	"database/sql/driver"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

const placeholder = "?"

type stmt struct {
	c    *conn
	q    []string
	argc int
}

//TODO: error handling
func newStmt(q string, c *conn) *stmt {
	q = strings.TrimRight(q, ";")
	st := new(stmt)
	r := strings.NewReader(q)
	strbuf := new(strings.Builder)
	var escape, arg bool
	for {
		ch, size, err := r.ReadRune()
		if size == 1 && !escape {
			switch ch {
			case '\\':
				escape = true
				continue
			case '?':
				arg = true
				st.q = append(st.q, strbuf.String())
				strbuf = new(strings.Builder)
				st.q = append(st.q, placeholder)

			default:
			}
		}
		if size > 0 && !arg {
			escape = false
			strbuf.WriteRune(ch)
		}

		arg = false

		if err == io.EOF {
			if strbuf.Len() > 0 {
				st.q = append(st.q, strbuf.String())
			}
			break
		} else if err != nil {
			break
		}
	}

	st.c = c

	return st
}

func (s *stmt) Close() error {
	//noop
	s.q = nil
	return nil
}

func (s *stmt) qBuiler(args []driver.Value) (string, error) {
	if len(args) < s.NumInput() {
		return "", ErrStmtInvalidArgc
	}

	sb := new(strings.Builder)

	i := 0
	for k := range s.q {
		if s.q[k] == placeholder {
			switch args[i].(type) {
			case int64:
				sb.WriteString(strconv.FormatInt(args[i].(int64), 10))
			case float64:
				sb.WriteString(strconv.FormatFloat(args[i].(float64), 'f', -1, 32))
			case bool:
				if args[i].(bool) {
					sb.WriteString("true")
				} else {
					sb.WriteString("false")
				}
			case []byte:
			case string:
				sb.WriteByte('`')
				sb.WriteString(args[i].(string))
				sb.WriteByte('`')
			case time.Time:
				sb.WriteString("``")
			}

			i++
		} else {
			sb.WriteString(s.q[k])
		}
	}
	return sb.String(), nil
}

//TODO: refactor
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {

	q, err := s.qBuiler(args)
	if err != nil {
		return nil, err
	}

	c := s.c.c.Cursor()
	c.Exec(context.TODO(), q)

	if err := c.Error(); err != nil {
		return &result{0, 0, err}, err
	}

	return &result{0, 0, nil}, nil
}

func (s *stmt) NumInput() int {
	if s.argc == 0 {
		for i := range s.q {
			if s.q[i] == placeholder {
				if s.argc == math.MaxInt32 {
					panic("argc overflow")
				}
				s.argc++
			}
		}
	}
	return s.argc
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	q, err := s.qBuiler(args)
	if err != nil {
		return nil, err
	}
	c := s.c.c.Cursor()
	c.Exec(context.TODO(), q)

	if err = c.Error(); err == nil {
		r := newRow(context.TODO(), c)
		return r, nil
	}

	return nil, err
}

// // ExecContext implements driver.StmtExecContext
// func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
// 	return nil, nil
// }

// // QueryContext implements driver.StmtQueryContext
// func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
// 	return nil, nil
// }
