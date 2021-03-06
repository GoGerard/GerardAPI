# GerardAPI
A modern bot for Discord - The Web API.

[![Build Status](https://travis-ci.org/GoGerard/GerardAPI.svg)](https://travis-ci.org/GoGerard/GerardAPI)
[![Go Report](http://goreportcard.com/badge/GoGerard/GerardAPI)](http://goreportcard.com/report/GoGerard/GerardAPI)
[![Issues](https://img.shields.io/github/issues/GoGerard/GerardAPI.svg)](https://github.com/GoGerard/GerardAPI/issues)
[![Coverage Status](https://coveralls.io/repos/GoGerard/GerardAPI/badge.svg?branch=master&service=github)](https://coveralls.io/github/GoGerard/GerardAPI?branch=master)

----------

## GoGerard

Project details can be found on the [main repo.](https://github.com/GoGerard/GoGerard)

----------

### GerardAPI ###

The API server is written in Golang and is currently depended on the following libraries:

 - [Gorm](https://github.com/jinzhu/gorm) - 'The fantastic ORM library for Golang, aims to be developer friendly.'
 - [Gin](https://github.com/gin-gonic/gin) - 'Gin is a HTTP web framework written in Go (Golang). '


**How to use?**

 1. Clone repo
 2. Make new Postgres contrainer, or use shared container that is also used by GerardDiscord.
 3. `$ docker run --name POSTGRESCONTAINERNAME -e POSTGRES_PASSWORD=mysecretpassword -d postgres`
 4. Open terminal in project folder
 5. `$ docker build -t gerardapi .`
 6. `$ docker run -it --rm --name containername --link POSTGRESCONTAINERNAME:postgres -p 8080:8080 gerardapi`
