https://gist.github.com/progrium/1a6886e9b4d76ab6e76f


Library for defining configuration based on environment variables. Sort of a cross
between `kelseyhightower/envconfig` and the `flag` package. Example usage:

```
import "github.com/gliderlabs/envconfig"

// Basic types
var hostIp = envconfig.String("registrator_ip", "",
  "IP for ports mapped to the host")
var internal = envconfig.Bool("registrator_internal", false,
  "Use internal ports instead of published ones")
var ttlRefresh = envconfig.Int("registrator_ttl_refresh", 0,
  "Frequency with which service TTLs are refreshed")
var velocity = envconfig.Float64("velocity", 1.23,
  "Floating point value with no prefix")

// Convenience + validation
var deregister = envconfig.StringOption("registrator_deregister", "always",
  []string{"always", "never", "on-success"},
  "Deregister mode")
var hosts = envconfig.Strings("registrator_hosts", nil, ",", "List of hosts")
```

### Notes

 * The above functions return typed values
 * `nil` as default value for `Strings` means empty string slice
 * There are also Int64, Uint, and Uint64 types.
 * No pointer values. Environment is read and value is set immediately. No `envconfig.Parse()`
 * Prefix and variable name are normalized to uppercase when reading environ
 * Bool can have these case insensitive values:
   * true: true, 1, y
   * false: false, 0, n, "" (empty string)
 * Internally, each var is registered into a global map registry with struct something like:

```
ConfigVar {
  type
  name
  prefix
  default
  description
  value
}
```

 * These are exposed as:

```
 envconfig.Var(name) ConfigVar
 envconfig.Vars() []ConfigVar
```
