package extnet

import (
	"net"
	"strings"
)

// IsErrClosed is error closed
func IsErrClosed(err error) bool {
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

// IsErrTimeout is net error timeout
func IsErrTimeout(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}

// IsErrTemporary is net error timeout
func IsErrTemporary(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(net.Error)
	return ok && e.Temporary()
}

// IsErrRefused is error connection refused
func IsErrRefused(err error) bool {
	return err != nil && strings.Contains(err.Error(), "connection refused")
}

// IsErrDeadline is error i/o deadline reached
func IsErrDeadline(err error) bool {
	return err != nil && strings.Contains(err.Error(), "i/o deadline reached")
}

// IsErrSocketNotConnected is error socket is not connected
func IsErrSocketNotConnected(err error) bool {
	return err != nil && strings.Contains(err.Error(), "socket is not connected")
}
