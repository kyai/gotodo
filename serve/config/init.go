package config

import (
	"errors"
	"gotodo/serve/core"
	"gotodo/serve/flags"
	"regexp"
	"sync"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(flags.Data)
	viper.SetConfigName("todo")
	viper.SetConfigType("ini")
	viper.SafeWriteConfig()

	core.Bomb(viper.ReadInConfig())
}

type Key struct {
	name  string
	value string
	regex string
}

func (k Key) Get() string {
	if viper.IsSet(k.name) {
		return viper.GetString(k.name)
	}
	return k.value
}

func (k Key) Set(v string) error {
	if k.regex != "" && !regexp.MustCompile(k.regex).MatchString(v) {
		return errors.New("配置值无效")
	}
	viper.Set(k.name, v)
	return viper.WriteConfig()
}

var (
	keyMap   = make(map[string]*Key)
	keyMutex sync.Mutex
	keys     = make([]*Key, 0)
)

func key(name, value, regex string) *Key {
	k := &Key{name, value, regex}
	keyMap[name] = k
	keys = append(keys, k)
	return k
}

func GetKey(name string) *Key {
	keyMutex.Lock()
	defer keyMutex.Unlock()
	return keyMap[name]
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}
