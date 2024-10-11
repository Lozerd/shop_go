#!/bin/bash

echo "Waiting for PostgreSQL..."

while ! </dev/tcp/db/5432; do
  echo "..."
  sleep 0.1
done 2>/dev/null

echo "PostgreSQL started"

echo "Run application"
exec "$@"
