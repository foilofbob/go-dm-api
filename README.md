# go-dm-api

This a Golang API intended as a backend for a web app that manages a D&D note. 

The frontend code is located at [cljs-dm-client](https://github.com/foilofbob/cljs-dm-client)

## Setup
### Windows

1. Install golang
```bash
$ mise install
```

2. Install mysql [here](https://dev.mysql.com/downloads/mysql/)

### Mac
1. Install Brew
```bash
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
2. Install mysql
```bash
$ brew install mysql
```
3. Start mysql
```bash
$ brew services start mysql
```
4. Install mise
```bash
$ brew install mise
```
5. Install golang
```bash
$ mise install
```

### Everyone

Install gomigrate

```bash
$ mise run install-gomigrate
```


In the `config` directory, copy `config.yml.example` as `config.yml`. 
You will need to supply the password for your local DB, it is expecting to use the user `dm_tool` (configurable) 
and the database `dm_campaign_manager`.

### Initialize DB
First you'll need to create the DB and User:
```
create database dm_campaign_manager;
create user 'dm_tool'@'localhost' identified by 'YOUR_PASSWORD_HERE';
grant all on dm_campaign_manager.* to 'dm_tool'@'localhost';
```
Whatever password you choose, be sure it matches your `config.yml`

Create a new migration with `mise run create-migration add_location_tables`

Apply or revert all (or X number) migrations with
```
mise run migrate-up --pw pass1234
mise run migrate-down --pw pass1234 --x 2
```
**NOTE**: There are Mac specific task definitions for running migrations:
```
mise run mac-migrate-up --pw pass1234
mise run mac-migrate-down --pw pass1234 --x 2
```

## Running the API

We're using [air](https://github.com/air-verse/air) for hot reloading, run the app with: `mise run`

Or don't: `go run main.go`

Exercise it with curl: `curl http://localhost:8090/note/1` (or `curl.exe http://localhost:8090/note/1` on Windows)
