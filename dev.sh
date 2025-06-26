#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is running
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        print_error "Docker is not running. Please start Docker and try again."
        exit 1
    fi
}

# Check if .env file exists
check_env() {
    if [ ! -f .env ]; then
        print_error ".env file not found. Please create it with the required environment variables."
        exit 1
    fi
}

# Function to start development environment
start_dev() {
    print_status "Starting local development environment..."
    
    check_docker
    check_env
    
    print_status "Starting PostgreSQL..."
    make local-up
    
    print_success "PostgreSQL is ready!"
    print_status "Starting evms-indexer..."
    
    # Run the application
    make local-run
}

# Function to stop development environment
stop_dev() {
    print_status "Stopping local development environment..."
    make local-clean
    print_success "Development environment stopped!"
}

# Function to show logs
show_logs() {
    print_status "Showing PostgreSQL logs..."
    make logs
}

# Function to open database shell
db_shell() {
    print_status "Opening PostgreSQL shell..."
    make db-shell
}

# Function to run tests
run_tests() {
    print_status "Running tests..."
    make test
}

# Function to show help
show_help() {
    echo "evms-indexer Development Script"
    echo ""
    echo "Usage: $0 [COMMAND]"
    echo ""
    echo "Commands:"
    echo "  start     Start the development environment (PostgreSQL + application)"
    echo "  stop      Stop the development environment"
    echo "  logs      Show PostgreSQL logs"
    echo "  db        Open PostgreSQL shell"
    echo "  test      Run tests"
    echo "  help      Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 start    # Start development environment"
    echo "  $0 stop     # Stop development environment"
    echo "  $0 logs     # Show database logs"
}

# Main script logic
case "${1:-start}" in
    start)
        start_dev
        ;;
    stop)
        stop_dev
        ;;
    logs)
        show_logs
        ;;
    db)
        db_shell
        ;;
    test)
        run_tests
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        show_help
        exit 1
        ;;
esac 