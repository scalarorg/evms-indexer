package db

import (
	"fmt"

	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/evms-indexer/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbAdapter *DatabaseAdapter

// DatabaseAdapter implements btc.DBAdapter interface for reorg handling
type DatabaseAdapter struct {
	PostgresClient *gorm.DB
}

func NewDatabaseAdapter(connection string) (*DatabaseAdapter, error) {
	if dbAdapter == nil {
		pgClient, err := NewPostgresClient(connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres client: %w", err)
		}
		dbAdapter = &DatabaseAdapter{
			PostgresClient: pgClient,
		}
	}
	return dbAdapter, nil
}

func NewPostgresClient(connection string) (*gorm.DB, error) {
	if connection == "" {
		return nil, fmt.Errorf("connnection string is empty")
	}

	db, err := SetupDatabase(connection)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	return db, nil
}

func SetupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = RunMigrations(db)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	// Convert tables to hyper tables
	// if err := InitHyperTables(db); err != nil {
	// 	return nil, fmt.Errorf("failed to initialize hyper tables: %w", err)
	//}

	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&chains.BlockHeader{},
		&chains.BtcBlockHeader{},
		&chains.VaultTransaction{},
		&chains.TokenSent{},
		&chains.CommandExecuted{},
		&chains.ContractCall{},
		&chains.ContractCallApproved{},
		&chains.ContractCallWithToken{},
		&chains.TokenDeployed{},
		&chains.EvmRedeemTx{},
		&chains.SwitchedPhase{},
		&types.LogEventCheckPoint{},
	)
	if err != nil {
		return err
	}

	return nil
}
