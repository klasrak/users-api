#!/usr/bin/env sh
set -ex

arg="$1"

HOST="${arg%%:*}"
PORT="${arg#*:}"

until docker-compose exec postgres sh -c 'pg_isready'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
sleep 2
>&2 echo "Postgres is up - executing command"
