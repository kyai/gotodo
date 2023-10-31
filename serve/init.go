package serve

import (
	"gotodo/serve/core"
	"gotodo/serve/flags"
	"os"
)

func init() {
	// 创建数据目录
	if fi, err := os.Stat(flags.Data); err != nil || !fi.IsDir() {
		core.Bomb(os.MkdirAll(flags.Data, os.ModePerm))
	}

	// 创建APPKEY
	appKeyPath := flags.Data + "/todo.key"
	if _, err := os.Stat(appKeyPath); err != nil {
		core.Bomb(os.WriteFile(appKeyPath, []byte(core.RandString(32)), os.ModePerm))
	}

	// 读取APPKEY
	appKey, err := os.ReadFile(appKeyPath)
	core.Bomb(err)
	core.AppKey = string(appKey)
}
