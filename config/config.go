// Author: Lutong.li
package config

import (
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64
	UserName   string
	Password   string

	//incluster配置
	MetadataNamespace string //元数据存储的namespace
	//InclusterKubeconfigPath string // incluster kubeconfig 文件路径

	//incluster clientset
	InclusterClientSet *kubernetes.Clientset
)

func init() {
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("JWT_SIGN_KEY", "lutong")

	viper.SetDefault("USERNAME", "lutong")
	viper.SetDefault("PASSWORD", "lutong123123")

	viper.SetDefault("METADATA_NAMESPACE", "TRM")
	MetadataNamespace = viper.GetString("METADATA_NAMESPACE")

	//logLevel := viper.GetString("LOG_LEVEL")

	//incluster相关配置
	viper.SetDefault("METADATA_NAMESPACE", "krm")
	//viper.SetDefault("INCLUSTER_KUBE_CONFIG_PATH", "~/.kube/config")

	//读取配置文件
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXP_TIME")
	UserName = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	MetadataNamespace = viper.GetString("METADATA_NAMESPACE")

	//日志的相关配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()

	//自动将环境变量映射到配置键
	viper.AutomaticEnv()
}

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadDefault() {
	c.Port = ":8080"
}
