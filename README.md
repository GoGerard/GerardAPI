# GerardAPI
A modern bot for Discord - The Web API.

[![Build Status](https://travis-ci.org/GoGerard/GerardAPI.svg)](https://travis-ci.org/GoGerard/GerardAPI)
[![Go Report](http://goreportcard.com/badge/GoGerard/GerardAPI)](http://goreportcard.com/report/GoGerard/GerardAPI)
[![Issues](https://img.shields.io/github/issues/GoGerard/GerardAPI.svg)](https://github.com/GoGerard/GerardAPI/issues)
[![Coverage Status](https://coveralls.io/repos/GoGerard/GerardAPI/badge.svg?branch=master&service=github)](https://coveralls.io/github/GoGerard/GerardAPI?branch=master)

----------

## GoGerard
GoGerard is an opensource project that focuses on easy to adapt, community-driven chatbots for [Discord](https://discordapp.com/).

The application is written in three separated parts, which are all replaceable to adapt to your project's needs.

 - [GerardDiscord](https://github.com/GoGerard/GerardDiscord) - A client, written in Golang, that communicates with the Discord API.
 - [GerardAPI](https://github.com/GoGerard/GerardAPI) - An API Server, written in Golang, that is used to communicate with the client and database(s)
 - [GerardJS](https://github.com/GoGerard/GerardJS) - A web interface, powered by AngularJS, that serves the API to its end-users.

Note that in the current state the project is nowhere finished, dependent on unstable external libraries,  and breakable changes to the project will happen till a future release.

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
