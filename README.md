# evms-indexer

Indexer for scalar transactions on EVMs and Bitcoin (via Electrum)

## Features

- **EVM Indexing**: Indexes events from multiple EVM chains
- **Electrum Indexing**: Direct Bitcoin block and VaultTransaction indexing via Electrum servers
- **Database Isolation**: Separate database connections for different components
- **Real-time Processing**: Processes blocks and events as they arrive
- **Recovery**: Automatic recovery of missing events and blocks

## Local Development

### Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Make

### Quick Start

1. **Start PostgreSQL with Docker Compose:**

   ```bash
   make local-up
   ```

2. **Run the application locally:**

   ```bash
   make local-run
   ```

3. **Or do both in one command:**

   ```bash
   make local-dev
   ```

### Development Commands

- `make local-up` - Start PostgreSQL database
- `make local-run` - Run the application with `go run`
- `make local-dev` - Start database and run application
- `make local-down` - Stop PostgreSQL database
- `make local-clean` - Stop database and clean up volumes

### Configuration

The application uses the following environment variables (defined in `.env`):

- `POSTGRES_USER` - PostgreSQL username (default: evms_indexer)
- `POSTGRES_PASSWORD` - PostgreSQL password (default: evms_indexer_password)
- `POSTGRES_DB` - PostgreSQL database name (default: evms_indexer)
- `POSTGRES_PORT` - PostgreSQL port (default: 5432)
- `CONFIG_PATH` - Path to configuration files (default: ./example/config)
- `DATABASE_URL` - Full database connection string
- `LOG_LEVEL` - Logging level (default: debug)

### Configuration Files

#### EVM Networks

EVM network configurations are stored in `example/config/evm.json`. Update this file to add or modify EVM networks.

#### Electrum Indexer

Electrum indexer configurations are stored in `example/config/electrum-indexer.json`. This configures Bitcoin block indexing via Electrum servers.

#### Separate Database Configuration

For separate database connections, use:

- `example/config/evm-separate-db.json` - EVM chains with individual databases
- Each component can have its own database connection for isolation

### Database Setup

#### Shared Database (Default)

The application automatically creates the necessary database tables and indexes on startup using a shared database.

#### Separate Databases

For production deployments, you can configure separate databases:

```sql
-- Electrum database
CREATE DATABASE electrum_db;
CREATE USER electrum_user WITH PASSWORD 'electrum_password';
GRANT ALL PRIVILEGES ON DATABASE electrum_db TO electrum_user;

-- EVM chain databases
CREATE DATABASE evm_sepolia_db;
CREATE USER evm_sepolia_user WITH PASSWORD 'evm_sepolia_password';
GRANT ALL PRIVILEGES ON DATABASE evm_sepolia_db TO evm_sepolia_user;
```

### Architecture

The indexer consists of multiple components:

1. **EVM Clients**: Index events from EVM chains
2. **Electrum Indexers**: Index Bitcoin blocks and VaultTransactions
3. **Database Adapters**: Handle data persistence
4. **Recovery Systems**: Recover missing events and blocks

### Stopping the Application

Use `Ctrl+C` to gracefully stop the application, or run:

```bash
make local-clean
```

## Production

### Build

```bash
make build
```

### Docker

```bash
make docker-image
make docker-up
```

## Documentation

For detailed information about the Electrum Indexer, see [docs/electrum-indexer.md](docs/electrum-indexer.md).
