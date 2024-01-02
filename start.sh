#!/bin/sh

set -e

echo "Running db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting the app"
exec "$@"
