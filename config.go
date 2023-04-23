package config

import (
	"time"

	"github.com/spf13/cast"
)

// StrMap 简写 —— map[string]interface{}
type StrMap map[string]interface{}

// Reload 重新加载配置
func Reload() error {
	return viperInstance.ReadInConfig()
}

// Add 新增配置项
func Add(name string, configuration map[string]interface{}) {
	viperInstance.Set(name, configuration)
}

// Set 设置配置项目
func Set(name string, value interface{}) {
	viperInstance.Set(name, value)
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况
	if !viperInstance.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viperInstance.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt32(key string, defaultValue ...interface{}) int32 {
	return cast.ToInt32(Get(key, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetUint16(key string, defaultValue ...interface{}) uint16 {
	return cast.ToUint16(Get(key, defaultValue...))
}

func GetUint32(key string, defaultValue ...interface{}) uint32 {
	return cast.ToUint32(Get(key, defaultValue...))
}
func GetUint64(key string, defaultValue ...interface{}) uint64 {
	return cast.ToUint64(Get(key, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}

func GetFloat64(key string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(Get(key, defaultValue...))
}

func GetTime(key string, defaultValue ...interface{}) time.Time {
	return cast.ToTime(Get(key, defaultValue...))
}

func GetDuration(key string, defaultValue ...interface{}) time.Duration {
	return cast.ToDuration(Get(key, defaultValue...))
}

func GetIntSlice(key string, defaultValue ...interface{}) []int {
	return cast.ToIntSlice(Get(key, defaultValue...))
}

func GetStringSlice(key string, defaultValue ...interface{}) []string {
	return cast.ToStringSlice(Get(key, defaultValue...))
}

func GetStringMap(key string, defaultValue ...interface{}) map[string]interface{} {
	return cast.ToStringMap(Get(key, defaultValue...))
}

func GetStringMapString(key string, defaultValue ...interface{}) map[string]string {
	return cast.ToStringMapString(Get(key, defaultValue...))
}

func GetStringMapStringSlice(key string, defaultValue ...interface{}) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key, defaultValue...))
}
