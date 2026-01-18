#!/bin/sh

set -e

# -----------------------------
# Validate prerequisites
# -----------------------------
if ! command -v docker >/dev/null 2>&1; then
  echo "❌ Docker is not installed"
  exit 1
fi

if ! command -v docker-compose >/dev/null 2>&1; then
  echo "❌ Docker Compose is not installed"
  exit 1
fi

# -----------------------------
# Create .env file
# -----------------------------
ENV_FILE=".env"

cat <<EOF > $ENV_FILE
# Server configuration
SERVER_ADDR=:8080
SERVER_WRITE_TIMEOUT=1m
SERVER_IDLE_TIMEOUT=5m
SERVER_READ_TIMEOUT=30s

# PostgreSQL configuration (Docker Compose)
POSTGRES_USER=postgres
POSTGRES_PASSWORD=admin
POSTGRES_SSL=disable
POSTGRES_PORT=5432
POSTGRES_DB=uber_coupons
POSTGRES_HOST=postgres
EOF

echo "✅ .env file created"

# -----------------------------
# Start Docker Compose
# -----------------------------
echo "🚀 Starting application using Docker Compose..."
docker-compose up --build

