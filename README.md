# evms-indexer

Indexer for scalar transactions on evms

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

EVM network configurations are stored in `example/config/evm.json`. Update this file to add or modify EVM networks.

### Database

The application automatically creates the necessary database tables and indexes on startup.

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
