#!/usr/bin/env bash

docker run -d --name relia-system \
  -e MYSQL_ROOT_PASSWORD="relia123456" \
  -e MYSQL_DATABASE="relia_system" \
  -e MYSQL_USER="relia" \
  -e MYSQL_PASSWORD="relia123456" \
  -e MYSQL_AUTHENTICATION_PLUGIN="mysql_native_password" \
  -p 3306:3306 \
  bitnami/mysql:8.0