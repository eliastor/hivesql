package hivesql

import (
	"reflect"
	"testing"
)

func Test_newStmt(t *testing.T) {
	type args struct {
		q string
		c *conn
	}
	tests := []struct {
		name string
		args args
		want *stmt
	}{
		{"no params", args{"SHOW DATABASES", nil},
			&stmt{c: nil, q: []string{"SHOW DATABASES"}},
		},
		{"no params with semicolon", args{"SHOW DATABASES;", nil},
			&stmt{c: nil, q: []string{"SHOW DATABASES"}},
		},
		{"string param", args{"SHOW TABLEs IN ?;", nil},
			&stmt{c: nil, q: []string{"SHOW TABLES IN ", "?"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStmt(tt.args.q, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
