package main

import (
	"flag"
	"fmt"
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

var pluginJsonFile string

func init() {
	flag.StringVar(&pluginJsonFile, "pluginFile", "plugin.json", "插件配置文件")
	flag.Parse()
}

func main() {

	err := build.BuildPluginSvr(pluginJsonFile)

	if err != nil {
		fmt.Println("BuildPluginSvr err", err)
	} else {
		fmt.Println("BuildPluginSvr success")
	}

}
