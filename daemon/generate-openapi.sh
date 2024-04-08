#!/bin/sh

set -eux

curl --fail -o ./frontend/src/backend/api/openapi.json ${SERVER_URL:-http://localhost:8080}/api/spec.json

docker-compose run --rm -u node -w /app frontend npm run openapi
