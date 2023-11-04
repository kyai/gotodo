package flags

import "flag"

var (
	Host string
	Port string
	Data string //数据目录
	View string //视图目录

	Debug bool //调试模式
)

func init() {
	flag.StringVar(&Host, "h", "0.0.0.0", "Http listen host")
	flag.StringVar(&Port, "p", "8000", "Http listen port")
	flag.StringVar(&Data, "d", "data", "Data files directory")
	flag.StringVar(&View, "v", "dist", "View files directory")
	flag.BoolVar(&Debug, "debug", false, "Debug mode")
	flag.Parse()
}
