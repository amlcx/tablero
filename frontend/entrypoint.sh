#!/bin/sh
set -e

# Wait for Postgres to be ready
echo "Waiting for Postgres..."
until nc -z postgres 5432; do
  sleep 1
done

echo "Running DB migrations..."
bun --bun run db:generate
bun --bun run db:migrate

# Start the app
exec "$@"