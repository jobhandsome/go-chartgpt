package config

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type FrontendYaml struct {
	Port  int         `yaml:"port"`
	Debug bool        `yaml:"debug"`
	MySQL MySQLConfig `yaml:"mysql"`
}

var Config *FrontendYaml
