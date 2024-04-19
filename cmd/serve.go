/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/shanipribadi/go-cookiecutter/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.With().Str("command", "server").Logger()
		defer func() {
			logger.Info().Msg("exiting")
			time.Sleep(time.Second)
		}()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cookieCutterService := &server.CookieCutterService{
		}
		srv := server.New(
			&server.ServerConfig{
				ListenAddress: viper.GetString("listen-address"),
				TlsListenAddress: viper.GetString("tls-listen-address"),
				TlsPrivateKey: viper.GetString("tls-private-key"),
				TlsCertificate: viper.GetString("tls-certificate"),
			},
			&server.ServerDependencies{
				Logger:       logger,
				CookieCutterService: cookieCutterService,
			},
		)

		g, gctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			signalChannel := make(chan os.Signal, 1)
			signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

			select {
			case sig := <-signalChannel:
				logger.Debug().Str("signal", sig.String()).Msg("Received signal")
				cancel()
			case <-gctx.Done():
				logger.Debug().Msg("closing signal goroutine")
				return gctx.Err()
			}

			return nil

		})
		g.Go(func() error {
			logger.Info().Msg("Starting server")
			return srv.Start(ctx)
		})

		return g.Wait()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("listen-address", ":8080", "listen address")
	serveCmd.Flags().String("tls-listen-address", ":8443", "tls listen address")
	serveCmd.Flags().String("tls-private-key", "", "path to private key")
	serveCmd.Flags().String("tls-certificate", "", "path to certificate")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
