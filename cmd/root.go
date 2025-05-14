package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/indexer"
	"github.com/scalarorg/evms-indexer/pkg/db"
	"github.com/spf13/cobra"
)

var (
	environment string
	configPath  string
	rootCmd     = &cobra.Command{
		Use:   "indexer",
		Short: "Indexer for EVMs",
		Run:   run,
	}
)

//	func init() {
//		rootCmd.Flags().StringVarP(&configPath, "config-path", "c", "/app/config", "Path to the config directory")
//	}
func Execute() error {

	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	// Initialize logger
	config.InitLogger()
	if err := config.LoadEnv(environment); err != nil {
		log.Warn().Msgf("Failed to load config: %v", err)
	}
	// Initialize global DatabaseClient
	dbAdapter, err := db.NewDatabaseAdapter(config.GlobalConfig.ConnnectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create database adapter")
	}
	// Initialize indexer service
	service, err := indexer.NewService(&config.GlobalConfig, dbAdapter)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create evms indexer service")
	}

	// Start relayer service
	err = service.Start(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start evms indexer service")
	}

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down evms indexer...")
	service.Stop()
}
