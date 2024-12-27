### What is this?

This a Golang API intended as a backend for a web app that manages a D&D campaign. 

The frontend code is located at [cljs-dm-client](https://github.com/foilofbob/cljs-dm-client)

### Setup
You will need `Golang (v1.23.4)` installed as well as `MySQL 8`.

In the `config` directory, copy `config.yml.example` as `config.yml`. 
You will need to supply the password for local DB, it is expecting to use the user `dm_tool` (configurable) 
and the database `dm_campaign_manager` (not currently configurable). 

TODO: Proper DB migrations

### Run the API

Start the API: `go run main.go`

Exercise it with curl: `curl http://localhost:8090/campaign/1`
