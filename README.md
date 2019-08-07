# gopkg
some golang package encapsulations or useful utils
## env
基于[viper](https://github.com/spf13/viper)的包装

### 加载配置文件
配置文件支持`JSON`, `TOML`, `YAML`, `HCL`, 和 `Java properties config files`
```go
// 不带后缀的配置文件名称
// 只有配置文件名称则使用当前路径
env.Load(configName, filePaths...)
```

### 读取配置文件里的值
使用
- Get(key string, defaultValue ...interface{}) : interface{} (是GetValue的别名)
- GetValue(key string, defaultValue ...interface{}) : interface{}

可以读取配置文件中的值。

可以提供一个默认值，当配置文件中没有提供`key`对应的值时，使用这个默认值。

也可以直接获取配置信息的具体类型(也支持设置默认值):
- GetBool(key string, defaultValue ...bool) : bool
- GetFloat64(key string, defaultValue ...float64) : float64
- GetInt(key string, defaultValue ...int) : int
- GetString(key string, defaultValue ...string) : string
- GetStringMap(key string, defaultValue ...map[string]interface{}) : map[string]interface{}
- GetStringMapString(key string, defaultValue ...map[string]string) : map[string]string
- GetStringSlice(key string, defaultValue ...[]string) : []string
- GetTime(key string, defaultValue ...time.Time) : time.Time
- GetDuration(key string, defaultValue ...time.Duration) : time.Duration

获取所有配置可以使用 `AllSettings() : map[string]interface{}`

### 例子
使用[toml](https://github.com/toml-lang/toml)格式的配置文件
```toml
# This is a TOML document.

title = "TOML Example"

[database]
server = "192.168.1.1"
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true
```

```go
GetString("title") // (returns "TOML Example")
GetString("database.server", "127.0.0.1") // (returns "192.168.1.1")
GetString("database.username", "root") // (returns "root")
GetString("database.password") // (returns "")
```