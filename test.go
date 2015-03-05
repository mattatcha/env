// generated code -- DO NOT EDIT
package envconfig

import "net"

func String(name string, defaultVal string, usage string) string {
	v := Var(newStringValue(defaultVal), name, usage)
	return v.Value.Get().(string)
}

func Bool(name string, defaultVal bool, usage string) bool {
	v := Var(newBoolValue(defaultVal), name, usage)
	return v.Value.Get().(bool)
}

func Float64(name string, defaultVal float64, usage string) float64 {
	v := Var(newFloat64Value(defaultVal), name, usage)
	return v.Value.Get().(float64)
}

func Int(name string, defaultVal int, usage string) int {
	v := Var(newIntValue(defaultVal), name, usage)
	return v.Value.Get().(int)
}

func Int64(name string, defaultVal int64, usage string) int64 {
	v := Var(newInt64Value(defaultVal), name, usage)
	return v.Value.Get().(int64)
}

func Uint(name string, defaultVal uint, usage string) uint {
	v := Var(newUintValue(defaultVal), name, usage)
	return v.Value.Get().(uint)
}

func Uint64(name string, defaultVal uint64, usage string) uint64 {
	v := Var(newUint64Value(defaultVal), name, usage)
	return v.Value.Get().(uint64)
}

func IP(name string, defaultVal net.IP, usage string) net.IP {
	v := Var(newIPValue(defaultVal), name, usage)
	return v.Value.Get().(net.IP)
}
