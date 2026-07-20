package main

import (
	"flag"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/update/config"
)

const (
	contentMd5Header = "Content-Md5"
	filename         = "goctl"
)

var configFile = flag.String("f", "etc/update-api.json", "the config file")

func forChksumHandler(file string, next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	fs := http.FileServer(http.Dir(c.FileDir))
	http.Handle(c.FilePath, http.StripPrefix(c.FilePath, forChksumHandler(path.Join(c.FileDir, filename), fs)))
	logx.Must(http.ListenAndServe(c.ListenOn, nil))
}
