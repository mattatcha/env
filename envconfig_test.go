package envconfig

import (
	"fmt"
	"os"
	"testing"
)

func boolString(s string) string {
	if s == "0" {
		return "false"
	}
	return "true"
}

func ResetForTesting() {
	DefaultEnv = NewEnvSet(os.Args[0])
}

func TestEverything(t *testing.T) {
	ResetForTesting()
	Bool("test_bool", false, "bool value")
	Int("test_int", 0, "int value")
	Int64("test_int64", 0, "int64 value")
	Uint("test_uint", 0, "uint value")
	Uint64("test_uint64", 0, "uint64 value")
	String("test_string", "0", "string value")
	StringOption("test_string_opt", "0", []string{"0", "2"}, "string opt")
	Float64("test_float64", 0, "float64 value")
	Duration("test_duration", 0, "time.Duration value")

	desired := "0"
	visitor := func(c *ConfigVar) {
		if len(c.Name) > 5 && c.Name[0:5] == "TEST_" {
			ok := false
			switch {
			case c.Value.String() == desired:
				ok = true
			case c.Name == "TEST_BOOL" && c.Value.String() == boolString(desired):
				ok = true
			case c.Name == "TEST_DURATION" && c.Value.String() == desired+"s":
				ok = true
			}
			if !ok {
				t.Error(c.Value.String(), c.Name)
			}
		}
	}
	VisitAll(visitor)
}

func TestEnvSet(t *testing.T) {
	ResetForTesting()
	String("test_string", "0", "string value")
	set := NewEnvSet("test")

	str := set.String("test_string_set", "hello", "string value")
	if str != "hello" {
		t.Error("env set string should be `hello`, is ", str)
	}

}

func TestString(t *testing.T) {
	ResetForTesting()
	s := String("conf_string", "foo", "test string")
	if s != "foo" {
		t.Errorf("expected: %s got: %s", "foo", s)
	}
}

func TestStringOption(t *testing.T) {
	ResetForTesting()
	s := StringOption("conf_string_opt", "foo", []string{"foo", "bar"}, "set value must be an option")
	if s != "foo" {
		t.Errorf("expected: %s got: %s", "foo", s)
	}
}

func TestStringOptionModify(t *testing.T) {
	ResetForTesting()
	os.Setenv("CONF_STRING_OPT", "bar")
	s := StringOption("conf_string_opt", "foo", []string{"foo", "bar"}, "set value must be an option")
	if s != "bar" {
		t.Errorf("expected: %s got: %s", "bar", s)
	}
}

func TestStringOptionInvalid(t *testing.T) {
	ResetForTesting()
	os.Setenv("CONF_STRING_OPT", "foobar")
	s := StringOption("conf_string_opt", "foo", []string{"foo", "bar"}, "set value must be an option")
	if s != "foo" {
		t.Errorf("expected: %s got: %s", "bar", s)
	}
}

func TestSecret(t *testing.T) {
	ResetForTesting()
	os.Setenv("CONF_SECRET", "12345678")
	s := Secret("conf_secret", "test secret")
	if s != "12345678" {
		t.Errorf("expected: %s got: %s", "12345678", s)
	}
	if fmt.Sprintf("%s", Var("conf_secret").Value) != "XXXX5678" {
		t.Errorf("expected: %s got: %s", "XXXX5678", s)
	}
}

func TestRace(t *testing.T) {
	ResetForTesting()
	done := make(chan struct{})
	go func() {
		s := String("conf_string_go", "foo", "")
		if s != "foo" {
			t.Errorf("expected: %s got: %s", "bar", s)
		}
		close(done)
	}()

	s := String("conf_string", "foo", "")
	if s != "foo" {
		t.Errorf("expected: %s got: %s", "bar", s)
	}

	<-done
}

func TestPrintDefaults(t *testing.T) {
	ResetForTesting()
	PrintDefaults(nil) // TODO: replace with buffer and actually test
}

func TestPrintEnv(t *testing.T) {
	ResetForTesting()
	PrintEnv(nil, false, false) // TODO: replace with buffer and actually test
}
