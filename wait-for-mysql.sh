#!/bin/bash

# wait-for-mysql.sh

HOST=$1
PORT=$2
TIMEOUT=$3

until mysqladmin ping -h "$HOST" --port="$PORT" --silent; do
    echo "Waiting for MySQL to be available..."
    sleep 2
done

echo "MySQL is ready!"
