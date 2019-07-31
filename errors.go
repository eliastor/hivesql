package hivesql

import (
	"errors"
)

//Error in package
var (
	ErrInvalidConn       = errors.New("invalid connection")
	ErrMalformPkt        = errors.New("malformed packet")
	ErrNoTLS             = errors.New("TLS requested but server does not support TLS")
	ErrCleartextPassword = errors.New("this user requires clear text authentication. If you still want to use it, please add 'allowCleartextPasswords=1' to your DSN")
	ErrNativePassword    = errors.New("this user requires mysql native password authentication")
	ErrOldPassword       = errors.New("" /* 191 byte string literal not displayed */)
	//ErrUnknownPlugin     = errors.New("this authentication plugin is not supported")
	//ErrOldProtocol       = errors.New("Hive server does not support required protocol 41")
	ErrPktSync               = errors.New("commands out of sync. You can't run this command now")
	ErrPktSyncMul            = errors.New("commands out of sync. Did you run multiple statements at once?")
	ErrPktTooLarge           = errors.New("packet for query is too large. Try adjusting the 'max_allowed_packet' variable on the server")
	ErrBusyBuffer            = errors.New("busy buffer")
	ErrInvalidDSNEmptyCAFile = errors.New("invalid DSN with empty CA file")
	ErrInvalidDSNTLSConfig   = errors.New("invalid DSN with unknow tls property value")
	ErrInvalidDSNUnescaped   = errors.New("Unsecaped DSN")

	errInvalidDSNAddr    = errors.New("Invalid address")
	errInvalidDSNNoSlash = errors.New("DSN without slash")

	ErrStmtInvalidArgc = errors.New("Invalid number of arguments for prepared statement")
)
