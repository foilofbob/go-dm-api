# go-dm-api

This a Golang API intended as a backend for a web app that manages a D&D note. 

The frontend code is located at [cljs-dm-client](https://github.com/foilofbob/cljs-dm-client)

## Setup
You will need `Golang (v1.23.4)` installed as well as `MySQL 9`.

In the `config` directory, copy `config.yml.example` as `config.yml`. 
You will need to supply the password for your local DB, it is expecting to use the user `dm_tool` (configurable) 
and the database `dm_campaign_manager`.

[DB Design](https://lucid.app/lucidchart/5ebc88df-03e9-4eed-84a1-30d29e689dc4/edit?page=0_0&invitationId=inv_aa194cb4-066d-442d-93d0-a7ae430b0870#) (link to Lucid)

## Running the API

Start the API: `go run main.go`

Exercise it with curl: `curl http://localhost:8090/note/1`

For hot reloading you can run it with: `air` (see [air](https://github.com/air-verse/air))

Install with: `go install github.com/air-verse/air@latest`
