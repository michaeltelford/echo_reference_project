#!/bin/bash

set -e

# Wait for the DB...
sleep 3

# Run any DB migrations
bin/migrate \
-database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
-path "db/migrations" \
up

# Execute the main command
exec $@
