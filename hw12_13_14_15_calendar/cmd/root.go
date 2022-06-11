package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/configs/config"
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
	rootCmd.PersistentFlags().StringVar(&Config.Server.Host, "server.Host", "127.0.0.1", "server host")
	rootCmd.PersistentFlags().StringVar(&Config.Server.Port, "server.Port", "8080", "server port")
	rootCmd.PersistentFlags().StringVar(&Config.Storage.DSN, "storage.dsn", "", "db dsn")
	cobra.CheckErr(viper.BindEnv("logger.loglevel", "loglevel", "log-level", "log_level"))
	cobra.CheckErr(viper.BindEnv("logger.logfile", "logfile", "log_file", "logFile"))
	cobra.CheckErr(viper.BindEnv("server.Host", "host", "Host", "server.host"))
	cobra.CheckErr(viper.BindEnv("server.Port", "server.port", "port", "Port"))
	cobra.CheckErr(viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")))
	cobra.CheckErr(viper.BindPFlag("logger.loglevel", rootCmd.PersistentFlags().Lookup("logger.loglevel")))
	cobra.CheckErr(viper.BindPFlag("logger.logfile", rootCmd.PersistentFlags().Lookup("logger.logfile")))
	cobra.CheckErr(viper.BindPFlag("server.Host", rootCmd.PersistentFlags().Lookup("server.Host")))
	cobra.CheckErr(viper.BindPFlag("server.Port", rootCmd.PersistentFlags().Lookup("server.Port")))
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
}
