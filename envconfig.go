package envconfig

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type EnvSet struct {
	sync.Mutex
	name string
	vars map[string]*ConfigVar
}

// ConfigVar represents a value from the environment.
type ConfigVar struct {
	Name        string
	Description string
	Value       Value  // value as set
	Default     string // default value (as text); for description message
	Secret      bool
}

// String retrieves a environment variable by name and parses it to a string
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) String(name string, defaultVal string, description string) string {
	v := e.NewVar(newStringValue(defaultVal), name, description)
	return v.Value.Get().(string)
}

// String retrieves a environment variable by name and parses it to a string
// defaultVal will be returned if the variable is not found.
func String(name string, defaultVal string, description string) string {
	return DefaultEnv.String(name, defaultVal, description)
}

// StringOption like String except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) StringOption(name string, defaultVal string, options []string, description string) string {
	v := e.NewVar(newStringValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(string) {
			return v.Value.Get().(string)
		}
	}
	return defaultVal
}

// StringOption like String except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func StringOption(name string, defaultVal string, options []string, description string) string {
	return DefaultEnv.StringOption(name, defaultVal, options, description)
}

// Secret retrieves a environment variable by name and parses it to a secret string
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Secret(name string, description string) string {
	v := e.NewVar(newSecretValue(""), name, description)
	v.Secret = true
	return v.Value.Get().(string)
}

// Secret retrieves a environment variable by name and parses it to a secret string
// defaultVal will be returned if the variable is not found.
func Secret(name string, description string) string {
	return DefaultEnv.Secret(name, description)
}

// Bool retrieves a environment variable by name and parses it to a bool
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Bool(name string, defaultVal bool, description string) bool {
	v := e.NewVar(newBoolValue(defaultVal), name, description)
	return v.Value.Get().(bool)
}

// Bool retrieves a environment variable by name and parses it to a bool
// defaultVal will be returned if the variable is not found.
func Bool(name string, defaultVal bool, description string) bool {
	return DefaultEnv.Bool(name, defaultVal, description)
}

// Float64 retrieves a environment variable by name and parses it to a float64
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Float64(name string, defaultVal float64, description string) float64 {
	v := e.NewVar(newFloat64Value(defaultVal), name, description)
	return v.Value.Get().(float64)
}

// Float64 retrieves a environment variable by name and parses it to a float64
// defaultVal will be returned if the variable is not found.
func Float64(name string, defaultVal float64, description string) float64 {
	return DefaultEnv.Float64(name, defaultVal, description)
}

// Float64Option like Float64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) Float64Option(name string, defaultVal float64, options []float64, description string) float64 {
	v := e.NewVar(newFloat64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(float64) {
			return v.Value.Get().(float64)
		}
	}
	return defaultVal
}

// Int retrieves a environment variable by name and parses it to a int
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Int(name string, defaultVal int, description string) int {
	v := e.NewVar(newIntValue(defaultVal), name, description)
	return v.Value.Get().(int)
}

// Int retrieves a environment variable by name and parses it to a int
// defaultVal will be returned if the variable is not found.
func Int(name string, defaultVal int, description string) int {
	return DefaultEnv.Int(name, defaultVal, description)
}

// IntOption like Int except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) IntOption(name string, defaultVal int, options []int, description string) int {
	v := e.NewVar(newIntValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(int) {
			return v.Value.Get().(int)
		}
	}
	return defaultVal
}

// Int64 retrieves a environment variable by name and parses it to a int64
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Int64(name string, defaultVal int64, description string) int64 {
	v := e.NewVar(newInt64Value(defaultVal), name, description)
	return v.Value.Get().(int64)
}

// Int64 retrieves a environment variable by name and parses it to a int64
// defaultVal will be returned if the variable is not found.
func Int64(name string, defaultVal int64, description string) int64 {
	return DefaultEnv.Int64(name, defaultVal, description)
}

// Int64Option like Int64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) Int64Option(name string, defaultVal int64, options []int64, description string) int64 {
	v := e.NewVar(newInt64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(int64) {
			return v.Value.Get().(int64)
		}
	}
	return defaultVal
}

// Uint retrieves a environment variable by name and parses it to a uint
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Uint(name string, defaultVal uint, description string) uint {
	v := e.NewVar(newUintValue(defaultVal), name, description)
	return v.Value.Get().(uint)
}

// Uint retrieves a environment variable by name and parses it to a uint
// defaultVal will be returned if the variable is not found.
func Uint(name string, defaultVal uint, description string) uint {
	return DefaultEnv.Uint(name, defaultVal, description)
}

// UintOption like Uint except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) UintOption(name string, defaultVal uint, options []uint, description string) uint {
	v := e.NewVar(newUintValue(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(uint) {
			return v.Value.Get().(uint)
		}
	}
	return defaultVal
}

// Uint64 retrieves a environment variable by name and parses it to a uint64
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) Uint64(name string, defaultVal uint64, description string) uint64 {
	v := e.NewVar(newUint64Value(defaultVal), name, description)
	return v.Value.Get().(uint64)
}

// Uint64 retrieves a environment variable by name and parses it to a uint64
// defaultVal will be returned if the variable is not found.
func Uint64(name string, defaultVal uint64, description string) uint64 {
	return DefaultEnv.Uint64(name, defaultVal, description)
}

// Uint64Option like Uint64 except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func (e *EnvSet) Uint64Option(name string, defaultVal uint64, options []uint64, description string) uint64 {
	v := e.NewVar(newUint64Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().(uint64) {
			return v.Value.Get().(uint64)
		}
	}
	return defaultVal
}

func (e *EnvSet) Duration(name string, defaultVal time.Duration, description string) time.Duration {
	v := e.NewVar(newDurationValue(defaultVal), name, description)
	return v.Value.Get().(time.Duration)
}

func Duration(name string, defaultVal time.Duration, description string) time.Duration {
	return DefaultEnv.Duration(name, defaultVal, description)
}

// IP retrieves a environment variable by name and parses it to a net.IP
// defaultVal will be returned if the variable is not found.
func (e *EnvSet) IP(name string, defaultVal net.IP, description string) net.IP {
	v := e.NewVar(newIPValue(defaultVal), name, description)
	return v.Value.Get().(net.IP)
}

// IP retrieves a environment variable by name and parses it to a net.IP
// defaultVal will be returned if the variable is not found.
func IP(name string, defaultVal net.IP, description string) net.IP {
	return DefaultEnv.IP(name, defaultVal, description)
}

func (e *EnvSet) VisitAll(fn func(*ConfigVar)) {
	for _, cfg := range e.vars {
		fn(cfg)
	}
}

func VisitAll(fn func(*ConfigVar)) {
	DefaultEnv.VisitAll(fn)
}

// NewVar retrieves a variable from the environment that is of type Value.
func (e *EnvSet) NewVar(value Value, name string, description string) *ConfigVar {
	e.Lock()
	defer e.Unlock()
	envVar := &ConfigVar{
		Name:        strings.ToUpper(name),
		Description: description,
		Value:       value,
		Default:     value.String(),
	}
	_, defined := e.vars[name]
	if defined {
		panic("env: " + name + " already defined.")
	}

	// This step is part of Parse() in flags pkg.
	if v := os.Getenv(envVar.Name); v != "" {
		envVar.Value.Set(v)
	}

	if e.vars == nil {
		e.vars = make(map[string]*ConfigVar)
	}
	e.vars[name] = envVar

	return envVar
}

// NewVar retrieves a variable from the environment that is of type Value.
func NewVar(value Value, name string, description string) *ConfigVar {
	return DefaultEnv.NewVar(value, name, description)
}

// Var retrieves a ConfigVar by name from the ConfigVar map.
func (e *EnvSet) Var(name string) *ConfigVar {
	e.Lock()
	defer e.Unlock()
	if v, ok := e.vars[name]; ok {
		return v
	}
	return nil
}

// Var retrieves a ConfigVar by name from the ConfigVar map.
func Var(name string) *ConfigVar {
	return DefaultEnv.Var(name)
}

// Vars retrieve all ConfigVars from the ConfigVar map.
func (e *EnvSet) Vars() map[string]*ConfigVar {
	e.Lock()
	defer e.Unlock()
	// TODO: this should return a copy of the map
	return e.vars
}

// Vars retrieve all ConfigVars from the ConfigVar map.
func Vars() map[string]*ConfigVar {
	return DefaultEnv.Vars()
}

// PrintDefaults prints the default values of all defined ConfigVars.
func (e *EnvSet) PrintDefaults(out io.Writer) {
	e.Lock()
	defer e.Unlock()
	// TODO: locking could be removed if this used Vars after copying is done
	for _, v := range e.vars {
		env := fmt.Sprintf("%s=%q", v.Name, v.Default)
		fmt.Fprintf(out, "%-30s # %s\n", env, v.Description)
	}
}

// PrintDefaults prints the default values of all defined ConfigVars.
func PrintDefaults(out io.Writer) {
	DefaultEnv.PrintDefaults(out)
}

// PrintEnv prints the set values of all defined ConfigVars.
func (e *EnvSet) PrintEnv(out io.Writer) {
	e.Lock()
	defer e.Unlock()
	// TODO: locking could be removed if this used Vars after copying is done
	for _, v := range e.vars {
		env := fmt.Sprintf("%s=%q", v.Name, v.Value)
		fmt.Fprintf(out, "%-30s # %s\n", env, v.Description)
	}
}

// PrintEnv prints the set values of all defined ConfigVars.
func PrintEnv(out io.Writer) {
	DefaultEnv.PrintEnv(out)
}

var DefaultEnv = NewEnvSet(os.Args[0])

func NewEnvSet(name string) *EnvSet {
	e := &EnvSet{
		name: name,
	}
	return e
}
