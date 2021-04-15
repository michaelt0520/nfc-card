# ONETAP
  ONETAP in the building

## Requirements
* Go 1.15.7
* Postgres 13

## Installation (Development setup)
OS X: (Catalina version 10.15.4)

- Install Homebrew
```sh
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

- Install Go
> Access
https://golang.org/doc/install
> Download and install

- Setup Postgresql
> You can install & config manually
```sh
brew install postgresql
brew services start postgresql
brew services enable postgresql
```

- Setup DB file
Clone `/configs/.env_sample.sh` to `/configs/.env.development.sh`
Add your local database connect information like `host`, `username`, `password` etc...

- DB create
```sh
sudo -u ${username} psql postgres
create database nfc_card_development;
exit
```

## Start

> Start server only:
```sh
make start
```
