package env

import (
	"path"
	"time"

	"github.com/spf13/viper"
)

func Load(configName string, filePaths ...string) error {
	viper.SetConfigName(path.Base(configName))
	if len(filePaths) > 0 {
		for _, filePath := range filePaths {
			viper.AddConfigPath(path.Dir(filePath))
		}
	} else {
		viper.AddConfigPath(path.Dir(configName))
	}

	return viper.ReadInConfig()
}

func Get(key string, defaultValue ...interface{}) interface{} {
	return GetValue(key, defaultValue...)
}

func GetValue(key string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.Get(key)
}

func GetBool(key string, defaultValue ...bool) bool {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetBool(key)
}

func GetFloat64(key string, defaultValue ...float64) float64 {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetFloat64(key)
}

func GetInt(key string, defaultValue ...int) int {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetInt(key)
}

func GetString(key string, defaultValue ...string) string {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetString(key)
}

func GetStringMap(key string, defaultValue ...map[string]interface{}) map[string]interface{} {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetStringMap(key)
}

func GetStringMapString(key string, defaultValue ...map[string]string) map[string]string {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetStringMapString(key)
}

func GetStringSlice(key string, defaultValue ...[]string) []string {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetStringSlice(key)
}

func GetTime(key string, defaultValue ...time.Time) time.Time {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetTime(key)
}

func GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	if len(defaultValue) > 0 {
		viper.SetDefault(key, defaultValue[0])
	}
	return viper.GetDuration(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.Unmarshal(rawVal, opts...)
}
