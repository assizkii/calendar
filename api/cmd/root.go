package cmd

import (
	"context"
	"fmt"
	"github.com/assizkii/calendar/api/internal/adapters/servers/grpc_server"
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"os"
	"time"
)

var client entities.EventServiceClient
var requestCtx context.Context
var requestOpts grpc.DialOption
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "A simple event calendar",
	Long:  `Event calendar`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/configs/conf.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func initClient() {
	// After Cobra root config init

	// Establish context to timeout if server does not respond
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	// Establish insecure grpc options (no TLS)
	requestOpts = grpc.WithInsecure()

	grpcPort := viper.GetString("host")

	client = grpc_server.StartClient(requestCtx, requestOpts, grpcPort)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".calendar" (without extension).
		viper.AddConfigPath("api/configs")
		viper.SetConfigName("conf")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
