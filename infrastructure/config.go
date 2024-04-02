package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	// // 设置配置文件的名字
	// viper.SetConfigName("dev")
	// // 设置文件的格式
	// viper.SetConfigType("yaml")
	// // 设置查找配置文件的路径为当前路径 . 表示项目的工作目录，也就是main.go同级的那个目录
	// viper.AddConfigPath("./config")

	// 设置配置文件的路径
	viper.SetConfigFile("./config/dev.yaml")

	// 读取配置文件中的数据到viper中
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 从viper中获取配置数据
	redisPort := viper.Get("redis.port")
	fmt.Printf("redisPort: %v %T \n", redisPort, redisPort) // 打印结果：redisPort: 6379 int

	redisPort = viper.GetString("redis.port")
	fmt.Printf("redisPort: %v %T \n", redisPort, redisPort) // 打印结果：redisPort: 6379 string

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	redisConfig := config.Redis
	fmt.Printf("redisConfig %+v\n", redisConfig) // 打印的结果：redisConfig {Host:localhost Port:6379 User: Password:abc123 DB:0}

	mysqlConfig := config.MySQL
	fmt.Printf("mysqlConfig %+v\n", mysqlConfig)

	// 先确定环境变量设置成功
	fmt.Printf("设置的环境变量: ENV=%v\n", os.Getenv("ENV"))

	// viper.SetDefault("env", "DEV")
	viper.BindEnv("env", "DEV")
	viper.AutomaticEnv()
	env := viper.Get("env")
	fmt.Println("env: ", env) // env:  PROD

	priority := config.Priority
	fmt.Printf("priority: %v \n", priority)

}

type Config struct {
	Redis    RedisConfig
	MySQL    MySQLConfig
	Priority string
}

type RedisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       int
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string `mapstructure:"db_name"`
}
