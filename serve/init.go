package serve

import (
	"gotodo/serve/core"
	"gotodo/serve/flags"
	"os"
)

func init() {
	// 创建数据目录
	core.Bomb(os.MkdirAll(flags.Data, os.ModePerm))

	// 创建APPKEY
	core.Bomb(os.WriteFile(flags.Data+"/todo.key", []byte(core.RandString(32)), os.ModePerm))

	// 读取APPKEY
	appKey, err := os.ReadFile(flags.Data + "/todo.key")
	core.Bomb(err)
	core.AppKey = string(appKey)
}
