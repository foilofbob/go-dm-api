# go-dm-api

This a Golang API intended as a backend for a web app that manages a D&D note. 

The frontend code is located at [cljs-dm-client](https://github.com/foilofbob/cljs-dm-client)

## Setup
You will need `Golang (v1.24.4)`, you can install this with `mise install`. Then you'll also need to install [MySQL 9](https://dev.mysql.com/downloads/mysql/).

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

Then you'll want to apply the db migrations. 
Migrations are run with golang-migrate, you can install this with `mise run install-gomigrate`.

Create a new migration with `mise run create-migration add_location_tables`

Apply or revert all (or X number) migrations with
```
mise run migrate-up --pw pass1234
mise run migrate-down --pw pass1234 --x 2
```

## Running the API

We're using [air](https://github.com/air-verse/air) for hot reloading, run the app with: `mise run`

Or don't: `go run main.go`

Exercise it with curl: `curl http://localhost:8090/note/1` (or `curl.exe http://localhost:8090/note/1` on Windows)
