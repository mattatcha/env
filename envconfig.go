package envconfig

import (
	"os"
	"strings"
)

var Environment = make(map[string]*ConfigVar)

type ConfigVar struct {
	Name        string
	Prefix      string
	Description string
	Value       Value  // value as set
	Default     string // default value (as text); for usage message
}

func Var(value Value, name string, usage string) *ConfigVar {
	envVar := &ConfigVar{
		Name:        strings.ToUpper(name),
		Description: usage,
		Value:       value,
		Default:     value.String(),
	}
	_, alreadythere := Environment[name]
	if alreadythere {
		panic("env: " + name + " already defined.")
	}

	actual := os.Getenv(envVar.Name)
	if actual != "" {
		envVar.Value.Set(actual)
	}

	Environment[name] = envVar

	return envVar
}

//
// func String(name, defaultVal, usage string) string {
// 	v := Var(newStringValue(defaultVal), name, usage)
// 	return v.Value.String()
// }
//
// func Bool(name string, defaultVal bool, usage string) bool {
// 	v := Var(newBoolValue(defaultVal), name, usage)
// 	return v.Value.Get().(bool)
// }
//
// func Int(name string, defaultVal int, usage string) int {
// 	v := Var(newIntValue(defaultVal), name, usage)
// 	return v.Value.Get().(int)
// }
//
// func Float64(name string, defaultVal float64, usage string) float64 {
// 	v := Var(newFloat64Value(defaultVal), name, usage)
// 	return v.Value.Get().(float64)
// }
