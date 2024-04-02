package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// 创建根命令，没有其他子命令的时候会执行这个命令
	rootCmd := &cobra.Command{
		// Use字段表示命令的名称
		Use: "mycli",
		// Short字段表示对命令的简单描述
		Short: "My CLI is a simple CLI application built with Cobra and Viper",
	}
	var message string
	// 添加一个持久标志，一个持久标志在整个CLI应用中保持一致
	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "A custom message")

	// 将一个标志绑定到viper配置
	viper.BindPFlag("message", rootCmd.PersistentFlags().Lookup("message"))
	// 设置默认的配置，当没有其他值使用的时候，就会使用默认配置的值
	viper.SetDefault("message", "Welcome to my CLI configured with Viper!")

	// 从配置文件中读取数据
	viper.SetConfigFile("./config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	priority := viper.Get("priority")
	fmt.Printf("priority: %v \n", priority)

	// 定义根命令运行的时候做些什么
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		message := viper.GetString("message")
		fmt.Printf("message: %v \n", message)
	}

	// 添加子命令
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of my cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("mycli v0.1")
		},
	}

	sayHelloCmd := &cobra.Command{
		Use:   "sayHello",
		Short: "Say Hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!")
		},
	}

	rootCmd.AddCommand(versionCmd, sayHelloCmd)

	// 执行cli应用
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
