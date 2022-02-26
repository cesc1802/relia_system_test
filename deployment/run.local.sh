#!/usr/bin/env bash

export JWT_SECRET_KEY="ABCD123456"
export LOG_DIR="./"
export LOG_FILENAME="log.log"
export SERVER_PORT="8088"
export DB_NAME="relia_system"
export DB_PASSWORD="relia123456"
export DB_HOST="localhost"
export DB_PORT="3306"
export DB_USER="relia"
export DB_CONN_TO="30"
export DB_READ_TO="30"
export DB_WRITE_TO="30"
export DB_MAX_CONN="50"
export DB_MAX_IDLE_CONN="50"
export DB_KEEP_ALIVE="30"

GO111MODULE=on  go build -o ../relia_system ../main.go
if [ $? -eq 0 ]
then
    cd ..
    ./relia_system
fi