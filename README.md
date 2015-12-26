# GerardAPI
A modern bot for Discord - The API


----------

## GoGerard
GoGerard is an opensource project that focuses on easy to adapt, community-driven chatbots for [Discord](https://discordapp.com/).

The application is written in three separated parts, which are all replaceable to adapt to your project's needs.

 - [GerardDiscord](https://github.com/GoGerard/GerardDiscord) - A client, written in Golang, that communicates with the Discord API. 
 - [GerardAPI](https://github.com/GoGerard/GerardAPI) - An API Server, written in Golang, that is used to communicate with the client and database(s)
 - [GerardJS](https://github.com/GoGerard/GerardJS) - A web interface, powered by AngularJS, that serves the API to its end-users. 

Note that in the current state the project is nowhere finished, dependent on unstable external libraries,  and breakable changes to the project will happen till a future release. 

The project will provide docker support in future.


----------

### GerardAPI ###

The API server is written in Golang and is currently depended on the following libraries: 

 - [Gorm](https://github.com/jinzhu/gorm) - 'The fantastic ORM library for Golang, aims to be developer friendly.'
 - [Gin](https://github.com/gin-gonic/gin) - 'Gin is a HTTP web framework written in Go (Golang). '

