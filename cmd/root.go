/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/KimMachineGun/automemlimit"
	_ "go.uber.org/automaxprocs"
)

var log zerolog.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cookiecutter",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())

		initLog()
		log.Debug().Str("log-level", viper.GetString("log-level")).Msg("rootCmd.PersistentPreRunE")
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func initLog() {
	wr := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
		// fmt.Printf("Logger Dropped %d messages", missed)
	})
	if viper.GetBool("log-unix-timestamp") {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}
  lvl, err := zerolog.ParseLevel(viper.GetString("log-level"))
  if err != nil {
    log = zerolog.New(wr).With().Timestamp().Logger()
  } else {
    log = zerolog.New(wr).Level(lvl).With().Timestamp().Logger()
  }
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().String("log-level", "debug", "log level")
	rootCmd.PersistentFlags().Bool("log-unix-timestamp", false, "log unix timestamp")
}
