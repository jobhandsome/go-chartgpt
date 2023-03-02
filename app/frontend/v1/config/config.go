package config

type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type Wechat struct {
	Appid  string `yaml:"appid"`
	Secret string `yaml:"secret"`
}

type FrontendYaml struct {
	Port   int    `yaml:"port"`
	Debug  bool   `yaml:"debug"`
	MySQL  Mysql  `yaml:"mysql"`
	Wechat Wechat `yaml:"wechat"`
}

var Config *FrontendYaml
