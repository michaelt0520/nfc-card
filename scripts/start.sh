#!/bin/bash

DEFAULT_ENV=development
DEFAULT_HOST=localhost
DEFAULT_PORT=8001

# check and set port for app
if [[ $port =~ ^[0-9]+$ ]]
then
  go_app_port=$port
else
  go_app_port=$DEFAULT_PORT
fi

# check and set env for app
if [[ "$env" = "production" ]]
then
  go_app_env=$env
else
  go_app_env=$DEFAULT_ENV
fi

# check and set host for app
if [[ "$host" != "" ]]
then
  go_app_host=$host
else
  go_app_host=$DEFAULT_HOST
fi

# export port and env
export app_port=$go_app_port
export app_env=$go_app_env
export app_host=$go_app_host

echo -e "\n\nONETAP run in: $app_env"
echo -e "ONETAP run: $app_host:$app_port\n\n"

# check existed or create file env
FILE="config/.env.${app_env}.sh"
if [ -f "$FILE" ]; then
  echo "$FILE exists"
else
  echo "$FILE does not exist"
  cp config/.env_sample.sh "config/.env.${app_env}.sh"
fi

go get -u go.uber.org/automaxprocs

# source env file
source ${FILE}

# check database existed or create database
if [ "$( psql -U $(whoami) postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$db_name'" )" = '1' ]
then
  echo "database $db_name already exists"
else
  psql -U $(whoami) postgres -c "create database $db_name;"
  echo "create database $db_name"
fi

# run app
go run main.go
