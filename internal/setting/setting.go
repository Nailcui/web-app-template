package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config *Configs

func init() {
	config := new(Configs)
	file, err := ioutil.ReadFile("configs/application.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	Config = config
}

type Configs struct {
	Mysql             *Mysql     `yaml:"mysql"`
	AppConfig         *App       `yaml:"appConfig"`
	JenkinsConfig     *Jenkins   `yaml:"jenkinsConfig"`
	CdiamondConfigMap *ConfigMap `yaml:"cdiamondConfigMap"`
	ApolloConfigMap   *ConfigMap `yaml:"apolloConfigMap"`
	Log               *Log       `yaml:"log"`
	Gitlib            *Gitlib    `yaml:"gitlib"`
}

type App struct {
	Host        string `yaml:"host"`
	CallbackURL string `yaml:"callbackURL"`
}

type Jenkins struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	JobToken string `yaml:"jobToken"`
}

type ConfigMap struct {
	Name      string            `yaml:"name"`
	MountPath string            `yaml:"mountPath"`
	Data      map[string]string `yaml:"data"`
	// todo BinaryData
	// BinaryData map[string][]byte
}

type Log struct {
	Path           string `yaml:"path"`
	Level          string `yaml:"level"`
	ConsoleEnabled bool   `yaml:"consoleEnabled"`
}

type Gitlib struct {
	Host         string `yaml:"host"`
	OAuthUrl     string `yaml:"OAuthUrl"`
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	GrantType    string `yaml:"grantType"`
	RedirectUri  string `yaml:"redirectUri"`
	LoginUrl     string `yaml:"loginUrl"`
}

type Mysql struct {
	Url         string `yaml:"url"`
	ShowLog     bool   `yaml:"showLog"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}
