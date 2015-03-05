package envconfig

import "testing"

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
