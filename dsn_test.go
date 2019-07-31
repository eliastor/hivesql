package hivesql

import (
	"reflect"
	"testing"
)

func Test_parseDSN(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *config
		wantErr bool
	}{
		{
			"simple DSN",
			args{"user:pass@localhost/"},
			&config{
				user:     "user",
				password: "pass",
				addr:     "localhost:10000",
				dbName:   "default",
			},
			false,
		},
		{
			"ip DSN",
			args{"user:pass@127.0.0.1/"},
			&config{
				user:     "user",
				password: "pass",
				addr:     "127.0.0.1:10000",
				dbName:   "default",
			},
			false,
		},
		{
			"db name",
			args{"user:pass@/blah"},
			&config{
				user:     "user",
				password: "pass",
				addr:     "localhost:10000",
				dbName:   "blah",
			},
			false,
		},
		{
			"hive opts",
			args{"user:pass@/?hive.opt1=123&hive.opt2=true"},
			&config{
				user:     "user",
				password: "pass",
				addr:     "localhost:10000",
				dbName:   "default",
				hiveConfig: map[string]string{
					"hive.opt1": "123",
					"hive.opt2": "true",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := parseDSN(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDSN() error = %v, wantErr %v, cfg %v", err, tt.wantErr, tt.wantCfg)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("parseDSN() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
