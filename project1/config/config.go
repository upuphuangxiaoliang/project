package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// 配置
type tomlConfig struct {
	// 定义结构映射
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动的时候 就会执行init方法
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	// 拿到路径     得到的路径是到D:\office\OneDrive\OneDrive - hkknv\code\Vscode\go\src\project\project1
	currenDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currenDir

	// 可以获取默认的数据  即/config.toml 里面已经给定的数据
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}

}
