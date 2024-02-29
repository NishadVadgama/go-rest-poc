# Audio GO - GO stack to support radio meta data microservice

# Minimum version requirements
- GO - v1.22
- Goose - v3.16.0

# Install Requirements
The following libraries will be needed (you will need [homebrew](https://brew.sh/)):

```bash
brew install go
brew install goose
```

## Local Setup

1. Copy the example environment file: `cp env.local.sh.example env.local.sh`
2. Modify `env.local.sh` as per your local setup
3. Source the environment file: `source env.local.sh`

### Run

```bash
go run main.go
```

### Migration

```bash
# To check migration status
goose --dir ./pkg/db/migration status

# To create a new migration
goose --dir ./pkg/db/migration create add_articles_table sql

# To push migration
goose --dir ./pkg/db/migration up
```

### Seeder

```bash
go run main.go -db_seed=true
```