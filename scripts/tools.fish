#!/usr/bin/fish

# Global variables
set -g MIGRATIONS_DIR migrations
set -g SQL_FILES_DIR sql
set -g DATABASE_PATH "$HOME/.local/share/.jadwal.db"
set -g DATABASE_URL "sqlite://$DATABASE_PATH?_journal_mode=WAL&_fk=1"
set -g DEV_URL "sqlite://:memory:"

cobra-cli completion fish | source

# Check if we're in a git repo and have sqlc.yml and go.mod
function _check_environment
    if not git rev-parse --is-inside-work-tree >/dev/null 2>&1
        echo "Error: Not inside a git repository"
        return 1
    end

    if not test -f "sqlc.yml"; or not test -f "go.mod"
        echo "Error: sqlc.yml or go.mod not found in current directory"
        return 1
    end

    if not command -v atlas >/dev/null 2>&1
        echo "Error: 'atlas' command not found in PATH"
        return 1
    end

    # Create migrations directory if it doesn't exist
    if not test -d "$MIGRATIONS_DIR"
        mkdir -p "$MIGRATIONS_DIR"
        echo "Created migrations directory: $MIGRATIONS_DIR"
    end

    return 0
end

# Generate migrations
function generate_migrations
    if not _check_environment
        return 1
    end

    echo "Generating migrations..."
    atlas migrate diff initial_schema \
        --dir "file://$MIGRATIONS_DIR" \
        --to "file://$SQL_FILES_DIR/schema.sql" \
        --dev-url "$DEV_URL"
end

# Apply migrations
function apply_migrations
    if not _check_environment
        return 1
    end

    echo "Applying migrations..."
    atlas migrate apply --url "$DATABASE_URL"
end

# Create empty migration
function empty_migration
    if not _check_environment
        return 1
    end

    if test -z "$argv[1]"
        echo "Usage: empty_migration <migration_name>"
        return 1
    end

    echo "Creating empty migration: $argv[1]"
    atlas migrate new "$argv[1]" --dir "file://$MIGRATIONS_DIR"
end

# Deactivate/unload the functions
function deactivate_tools
    set -e MIGRATIONS_DIR
    set -e SQL_FILES_DIR
    set -e DATABASE_PATH
    set -e DATABASE_URL
    set -e DEV_URL
    functions -e _check_environment generate_migrations apply_migrations empty_migration deactivate_migration_helpers
    echo "Migration helpers deactivated"
end

echo "Dev tools loaded. Available functions:"
echo "  generate_migrations"
echo "  apply_migrations"
echo "  empty_migration <name>"
echo "  deactivate_tools"
