package args

import (
	"flag"
)

type Args struct {
	Host string
	Port uint64
	Path string // 配置文件路径
}

func Parser() Args {
	var inputArgs Args
	flag.StringVar(&inputArgs.Host, "h", "localhost", "nacos host default: localhost")
	flag.Uint64Var(&inputArgs.Port, "p", 8848, "nacos port default: 8848")
	flag.StringVar(&inputArgs.Path, "c", "", "")
	flag.Parse()
	return inputArgs
}
