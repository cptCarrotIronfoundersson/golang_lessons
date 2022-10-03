package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/configs/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Config config.Config

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Calendar API SERVICE short annotation",
		Long:  `Calendar API SERVICE long annotation`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&Config.File.Path, "config", "../../configs/config.yaml",
		"Path to configuration file")
	rootCmd.PersistentFlags().StringVar(&Config.Logger.Level, "logger.loglevel", "INFO", "Logger level")
	rootCmd.PersistentFlags().StringVar(&Config.Logger.LogFile, "logger.logfile", "~/log", "Path to log File")
	rootCmd.PersistentFlags().StringVar(&Config.HTTPServer.Host, "http_server.Host", "127.0.0.1", "HTTP server host")
	rootCmd.PersistentFlags().StringVar(&Config.HTTPServer.Port, "http_server.Port", "8080", "GRPC SERVER port")
	rootCmd.PersistentFlags().StringVar(&Config.GRPCServer.Host, "grpc_server.Host", "127.0.0.1", "GRPC SERVER host")
	rootCmd.PersistentFlags().StringVar(&Config.GRPCServer.Port, "grpc_server.Port", "8081", "server port")
	rootCmd.PersistentFlags().StringVar(&Config.Storage.DSN, "storage.dsn", "", "db dsn")
	cobra.CheckErr(viper.BindEnv("logger.loglevel", "loglevel", "log-level", "log_level"))
	cobra.CheckErr(viper.BindEnv("logger.logfile", "logfile", "log_file", "logFile"))
	cobra.CheckErr(viper.BindEnv("http_server.Host", "httphost", "httpHost", "http_server.host"))
	cobra.CheckErr(viper.BindEnv("http_server.Port", "http_server.port", "httpPort"))
	cobra.CheckErr(viper.BindEnv("grpc_server.Host", "grpchost", "grpcHost", "grpc_server.host"))
	cobra.CheckErr(viper.BindEnv("grpc_server.Port", "grpc_server.port", "grpcPort"))
	cobra.CheckErr(viper.BindEnv("storage.dsn", "storageDsn"))
	cobra.CheckErr(viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")))
	cobra.CheckErr(viper.BindPFlag("storage.dsn", rootCmd.PersistentFlags().Lookup("storage.dsn")))
	cobra.CheckErr(viper.BindPFlag("logger.loglevel", rootCmd.PersistentFlags().Lookup("logger.loglevel")))
	cobra.CheckErr(viper.BindPFlag("logger.logfile", rootCmd.PersistentFlags().Lookup("logger.logfile")))
	cobra.CheckErr(viper.BindPFlag("http_server.Host", rootCmd.PersistentFlags().Lookup("http_server.Host")))
	cobra.CheckErr(viper.BindPFlag("http_server.Port", rootCmd.PersistentFlags().Lookup("http_server.Port")))
	cobra.CheckErr(viper.BindPFlag("grpc_server.Host", rootCmd.PersistentFlags().Lookup("grpc_server.Host")))
	cobra.CheckErr(viper.BindPFlag("grpc_server.Port", rootCmd.PersistentFlags().Lookup("grpc_server.Port")))
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if Config.File.Path != "" {
		viper.SetConfigFile(Config.File.Path)
	} else {
		// Find executable path
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		log.Printf("%s", err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		err := viper.Unmarshal(&Config)
		if err != nil {
			fmt.Println("Error during unmarshalling file:", err)
		}
	} else {
		cobra.CheckErr(err)
	}
	fmt.Println(Config.Storage.DSN)
}
