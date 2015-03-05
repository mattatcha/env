package envconfig

import (
	"os"
	"strings"
)

var Environment = make(map[string]*ConfigVar)

type ConfigVar struct {
	Name        string
	Description string
	Value       Value  // value as set
	Default     string // default value (as text); for description message
	Ext         map[string]interface{}
}

func newVar(value Value, name string, description string) *ConfigVar {
	envVar := &ConfigVar{
		Name:        strings.ToUpper(name),
		Description: description,
		Value:       value,
		Default:     value.String(),
	}
	_, exists := Environment[name]
	if exists {
		panic("env: " + name + " already defined.")
	}

	actual := os.Getenv(envVar.Name)
	if actual != "" {
		envVar.Value.Set(actual)
	}

	Environment[name] = envVar

	return envVar
}

func String(name, defaultVal, description string) string {
	v := newVar(newStringValue(defaultVal), name, description)
	return v.Value.String()
}

func Bool(name string, defaultVal bool, description string) bool {
	v := newVar(newBoolValue(defaultVal), name, description)
	return v.Value.Get().(bool)
}

func Int(name string, defaultVal int, description string) int {
	v := newVar(newIntValue(defaultVal), name, description)
	return v.Value.Get().(int)
}

func Float64(name string, defaultVal float64, description string) float64 {
	v := newVar(newFloat64Value(defaultVal), name, description)
	return v.Value.Get().(float64)
}
