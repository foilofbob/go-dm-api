# go-dm-api

This a Golang API intended as a backend for a web app that manages a D&D note. 

The frontend code is located at [cljs-dm-client](https://github.com/foilofbob/cljs-dm-client)

## Setup
You will need `Golang (v1.23.4)` installed as well as `MySQL 9`.

In the `config` directory, copy `config.yml.example` as `config.yml`. 
You will need to supply the password for your local DB, it is expecting to use the user `dm_tool` (configurable) 
and the database `dm_campaign_manager`.

[DB Design](https://lucid.app/lucidchart/5ebc88df-03e9-4eed-84a1-30d29e689dc4/edit?page=0_0&invitationId=inv_aa194cb4-066d-442d-93d0-a7ae430b0870#) (link to Lucid)

### Initialize DB
First you'll need to create the DB and User:
```
create database dm_campaign_manager;
create user 'dm_tool'@'localhost' identified by 'YOUR_PASSWORD_HERE';
grant all on dm_campaign_manager.* to 'dm_tool'@'localhost';
```
Whatever password you choose, be sure it matches your `config.yml`

Then you'll want to apply the db migrations. Migrations are run with golang-migrate.
```
go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create a new migration with
```
migrate create -ext sql -dir db/migrations -seq <title>
```

Apply or revert all (or X number) migrations with
```
migrate -path db/migrations -database "mysql://dm_tool:<password>@tcp(localhost:3306)/dm_campaign_manager" <up|down> <optional X>
```

## Running the API

Start the API: `go run main.go`

Exercise it with curl: `curl http://localhost:8090/note/1`

For hot reloading you can run it with: `air` (see [air](https://github.com/air-verse/air))

Install with: `go install github.com/air-verse/air@latest`
