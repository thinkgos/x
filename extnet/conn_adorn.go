package extnet

import (
	"compress/gzip"
	"compress/zlib"
	"crypto/tls"
	"net"

	"go.uber.org/atomic"

	"github.com/thinkgos/go-core-package/extnet/connection/cencrypt"
	"github.com/thinkgos/go-core-package/extnet/connection/cflow"
	"github.com/thinkgos/go-core-package/extnet/connection/cgzip"
	"github.com/thinkgos/go-core-package/extnet/connection/ciol"
	"github.com/thinkgos/go-core-package/extnet/connection/csnappy"
	"github.com/thinkgos/go-core-package/extnet/connection/czlib"
	"github.com/thinkgos/go-core-package/lib/encrypt"
)

// BaseAdornTLSClient base adorn tls client
func BaseAdornTLSClient(conf *tls.Config) func(conn net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		return tls.Client(conn, conf)
	}
}

// BaseAdornTLSServer base adorn tls server
func BaseAdornTLSServer(conf *tls.Config) func(conn net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		return tls.Server(conn, conf)
	}
}

// BaseAdornStcp base adorn encrypt with method and password
func BaseAdornStcp(method, password string) func(conn net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		cip, err := encrypt.NewCipher(method, password)
		if err != nil {
			panic("encrypt method should be valid")
		}
		return cencrypt.New(conn, cip)
	}
}

// AdornSnappy snappy chain
func AdornSnappy(compress bool) func(conn net.Conn) net.Conn {
	if compress {
		return func(conn net.Conn) net.Conn {
			return csnappy.New(conn)
		}
	}
	return func(conn net.Conn) net.Conn {
		return conn
	}
}

// AdornGzip gzip chain
func AdornGzip(compress bool) func(conn net.Conn) net.Conn {
	return AdornGzipLevel(compress, gzip.DefaultCompression)
}

// AdornGzipLevel gzip chain with level
// level see gzip package
func AdornGzipLevel(compress bool, level int) func(conn net.Conn) net.Conn {
	if compress {
		return func(conn net.Conn) net.Conn {
			return cgzip.NewLevel(conn, level)
		}
	}
	return func(conn net.Conn) net.Conn {
		return conn
	}
}

// AdornZlib zlib chain
func AdornZlib(compress bool) func(net.Conn) net.Conn {
	return AdornZlibLevel(compress, zlib.DefaultCompression)
}

// AdornZlibLevel zlib chain with the level
// level see zlib package
func AdornZlibLevel(compress bool, level int) func(net.Conn) net.Conn {
	return AdornZlibLevelDict(compress, level, nil)
}

// AdornZlibLevelDict zlib chain with the level and dict
// level see zlib package
func AdornZlibLevelDict(compress bool, level int, dict []byte) func(net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		if compress {
			return czlib.NewLevelDict(conn, level, dict)
		}
		return conn
	}
}

// AdornFlow cflow chain
func AdornFlow(Wc *atomic.Uint64, Rc *atomic.Uint64, Tc *atomic.Uint64) func(conn net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		return &cflow.Conn{Conn: conn, Wc: Wc, Rc: Rc, Tc: Tc}
	}
}

// AdornIol ciol chain
func AdornIol(opts ...ciol.Options) func(conn net.Conn) net.Conn {
	return func(conn net.Conn) net.Conn {
		return ciol.New(conn, opts...)
	}
}
