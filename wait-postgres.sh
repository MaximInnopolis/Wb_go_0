#!/bin/sh

set -e

host="postgres-wb-l0"
shift
cmd="$@"

until PGPASSWORD=${DBPASSWORD} psql -h "$host" -U ${DBUSERNAME} -d ${DBNAME} -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd