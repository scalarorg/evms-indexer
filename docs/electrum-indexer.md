# Electrum Indexer

The Electrum Indexer is a new component in the evms-indexer project that provides direct indexing of Bitcoin blocks and VaultTransactions without relying on external electrum libraries.

## Features

- **Direct RPC Client**: Uses direct JSON-RPC calls to electrum servers instead of importing external electrum libraries
- **Separate Database Connection**: Each electrum indexer has its own database connection for isolation
- **Block Header Indexing**: Indexes Bitcoin block headers with full metadata
- **VaultTransaction Parsing**: Parses and identifies ScalarVault transactions in blocks
- **Real-time Block Monitoring**: Continuously monitors for new blocks and processes them
- **Database Isolation**: Separate database connections for electrum and EVM data

## Architecture

### Components

1. **ElectrumIndexer**: Main indexer that connects to electrum servers
2. **ElectrumIndexerConfig**: Configuration for each indexer instance
3. **VaultTransaction**: Parsed vault transaction data structure
4. **BlockHeader**: Bitcoin block header information

### Database Schema

The electrum indexer creates its own tables:

- `block_headers`: Stores Bitcoin block header information
- `vault_transactions`: Stores parsed VaultTransaction data

## Configuration

### Electrum Indexer Configuration

Create a configuration file at `config/electrum-indexer.json`:

```json
[
  {
    "host": "localhost",
    "port": 50001,
    "database_url": "postgres://electrum_user:electrum_password@localhost:5432/electrum_db?sslmode=disable",
    "dial_timeout": "30s",
    "method_timeout": "60s",
    "ping_interval": "30s",
    "max_reconnect_attempts": 120,
    "reconnect_delay": "5s",
    "enable_auto_reconnect": true,
    "batch_size": 1,
    "confirmations": 1,
    "source_chain": "bitcoin-mainnet"
  }
]
```

### EVM Separate Database Configuration

For EVM clients with separate databases, use `config/evm-separate-db.json`:

```json
[
  {
    "chain_id": 11155111,
    "id": "ethereum-sepolia",
    "name": "Ethereum Sepolia",
    "rpc_url": "https://sepolia.infura.io/v3/YOUR_INFURA_KEY",
    "gateway": "0x1234567890123456789012345678901234567890",
    "finality": 12,
    "start_block": 0,
    "blockTime": 12000,
    "recover_range": 1000000,
    "database_url": "postgres://evm_sepolia_user:evm_sepolia_password@localhost:5432/evm_sepolia_db?sslmode=disable"
  }
]
```

## VaultTransaction Parsing

The indexer implements VaultTransaction parsing similar to the Rust code in the electrs project:

### Supported Versions

- **Version 1**: Basic vault transaction format
- **Version 2**: Enhanced vault transaction format
- **Version 3**: Full vault transaction format with all fields

### Parsing Logic

1. **OP_RETURN Detection**: Identifies transactions with OP_RETURN outputs
2. **SCALAR Tag**: Looks for the "SCALAR" service tag in OP_RETURN data
3. **Version Parsing**: Parses the version byte to determine format
4. **Field Extraction**: Extracts all relevant fields based on version

### Parsed Fields

- Transaction ID and block information
- Amount and addresses
- Service tag and covenant quorum
- Destination chain and token information
- Session sequence and custodian group UID
- Script pubkey and raw transaction data

## Usage

### Integration with Indexer Service

The electrum indexer is automatically integrated into the indexer service:

```go
// In indexer service
electrumIndexers, err := electrum.NewElectrumIndexers(config, dbAdapter)
if err != nil {
    return nil, fmt.Errorf("failed to create electrum indexers: %w", err)
}

// Start indexers
for _, indexer := range s.ElectrumIndexers {
    go func(idx *electrum.ElectrumIndexer) {
        err := idx.Start(ctx)
        if err != nil {
            log.Error().Err(err).Msg("Failed to start electrum indexer")
        }
    }(indexer)
}
```

### Database Operations

```go
// Get block header
blockHeader, err := indexer.GetBlockHeader(ctx, height)

// Index a block
err := indexer.IndexBlock(ctx, height)

// Get latest height
latestHeight, err := indexer.GetLatestHeight(ctx)
```

## Benefits

1. **No External Dependencies**: Doesn't import external electrum libraries, reducing dependency complexity
2. **Database Isolation**: Separate database connections prevent conflicts between electrum and EVM data
3. **Real-time Processing**: Processes blocks as they arrive
4. **Configurable**: Supports multiple electrum servers and configurations
5. **Robust**: Includes reconnection logic and error handling
6. **Scalable**: Each component can have its own database connection

## Monitoring

The indexer provides comprehensive logging:

- Connection status and reconnection attempts
- Block processing progress
- VaultTransaction parsing results
- Error conditions and recovery

## Database Setup

### Electrum Database

Create a separate database for electrum data:

```sql
CREATE DATABASE electrum_db;
CREATE USER electrum_user WITH PASSWORD 'electrum_password';
GRANT ALL PRIVILEGES ON DATABASE electrum_db TO electrum_user;
```

### EVM Separate Databases

For each EVM chain, create a separate database:

```sql
-- For Ethereum Sepolia
CREATE DATABASE evm_sepolia_db;
CREATE USER evm_sepolia_user WITH PASSWORD 'evm_sepolia_password';
GRANT ALL PRIVILEGES ON DATABASE evm_sepolia_db TO evm_sepolia_user;

-- For Polygon Mumbai
CREATE DATABASE evm_mumbai_db;
CREATE USER evm_mumbai_user WITH PASSWORD 'evm_mumbai_password';
GRANT ALL PRIVILEGES ON DATABASE evm_mumbai_db TO evm_mumbai_user;
```

## Future Enhancements

- Enhanced VaultTransaction parsing for all versions
- Support for additional electrum server features
- Performance optimizations for high-throughput scenarios
- Integration with additional blockchain networks
- Advanced reconnection strategies with exponential backoff
- Metrics and monitoring integration
