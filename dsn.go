package hivesql

import (
	"crypto/x509"
	"io/ioutil"
	"net/url"
	"strings"
)

func parseDSN(dsn string) (cfg *config, err error) {
	// New config with some default values
	cfg = new(config)

	// [user[:password]@][[proto://][addr]]/dbname[?param1=value1&paramN=valueN]
	// Find the last '/' (since the password or the net addr might contain a '/')
	foundSlash := false
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int

			// left part is empty if i <= 0
			if i > 0 {
				// [username[:password]@][protocol[(address)]]
				// Find the last '@' in dsn[:i]
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						// username[:password]
						// Find the first ':' in dsn[:j]
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								cfg.password = dsn[k+1 : j]
								break
							}
						}

						cfg.user = dsn[:k]

						break
					}
				}

				cfg.addr = dsn[j+1 : i]
			}

			// dbname[?param1=value1&...&paramN=valueN]
			// Find the first '?' in dsn[i+1:]
			for j = i + 1; j < len(dsn); j++ {
				if dsn[j] == '?' {
					if err = parseDSNParams(cfg, dsn[j+1:]); err != nil {
						return
					}
					break
				}
			}
			cfg.dbName = dsn[i+1 : j]

			break
		}
	}

	if !foundSlash && len(dsn) > 0 {
		return nil, errInvalidDSNNoSlash
	}

	cfg.normalize()

	return
}

func parseDSNParams(cfg *config, params string) (err error) {

	for _, v := range strings.Split(params, "&") {
		param := strings.SplitN(v, "=", 2)
		if len(param) != 2 {
			continue
		}

		// cfg params
		switch value := param[1]; param[0] {
		case "tlsCAFile":
			rootCAs := x509.NewCertPool()
			certs, err := ioutil.ReadFile(value)
			if err != nil {
				return err
			}
			if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
				return ErrInvalidDSNEmptyCAFile
			}
			cfg.turnTLS()
			cfg.tlsConfig.RootCAs = rootCAs

		case "ldap":
			boolValue, isBool := readBool(value)
			if isBool && boolValue {
				cfg.auth = "LDAP"
			} else {
				cfg.auth = "NONE"
			}

		// TLS-Encryption
		case "tls":
			cfg.turnTLS()
			boolValue, isBool := readBool(value)
			if isBool && boolValue {

			} else if strings.ToLower(value) == "skip-verify" {
				cfg.tlsConfig.InsecureSkipVerify = true
			} else {
				return
			}

		default:
			if cfg.hiveConfig == nil {
				cfg.hiveConfig = make(map[string]string)
			}

			if cfg.hiveConfig[param[0]], err = url.QueryUnescape(value); err != nil {
				return ErrInvalidDSNUnescaped
			}
		}
	}

	return
}

func readBool(input string) (value bool, valid bool) {
	switch input {
	case "1", "true", "TRUE", "True", "yes", "YES", "Yes":
		return true, true
	case "0", "false", "FALSE", "False", "no", "No", "NO":
		return false, true
	}

	// Not a valid bool value
	return
}
