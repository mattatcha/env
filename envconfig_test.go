package envconfig

import (
	"net"
	"testing"
)

func TestString(t *testing.T) {
	s := String("path", "d", "some test env var")
	t.Logf("%v", s)

	t.Log(Environment)
}

func TestBool(t *testing.T) {
	s := Bool("something", true, "some test env var")
	t.Logf("%v", s)

	t.Log(Environment)
}

func TestIP(t *testing.T) {
	s := IP("ipaddr", net.ParseIP("8.8.8.8"), "some test env var")
	t.Logf("%v", s)

	t.Log(Environment)
}
