package main

import (
	"flag"

	gc "github.com/sky0621/go-crud"
)

var Cfg *gc.Config

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	cfg := flag.String("config", "config.toml", "Config File")
	flag.Parse()

	gc.ReadConfig(*cfg)

	gc.Start()
}
