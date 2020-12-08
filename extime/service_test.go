package extime

import (
	"testing"
)

func TestServiceTime(t *testing.T) {
	t.Log(ServiceStartupTime())
	t.Log(ServiceElapseTime())
	t.Log(ServiceUptime())
}
