# Welcome to CrazyBaaS
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/crazycoderlabs/crazybaas.svg?branch=master)](https://travis-ci.org/crazycoderlabs/crazybaas)
[![GoReport](https://goreportcard.com/badge/github.com/crazycoderlabs/crazybaas)](https://goreportcard.com/report/github.com/crazycoderlabs/crazybaas)

CrazyBaaS is just another BaaS (Backend-as-as-Service) tool.

## Technology Stack
- golang
- Revel
## Architecture
- Plugin Based for Core
- Microservices for Wrappers

## Developer Notes

### Start the web server:

    revel run myapp

   Go to http://localhost:9000/ to login to the Admin UI

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Maintainers
`- Sarvesh Chitko (chitkosarvesh@gmail.com)`

*Made with :hearts: by {CrazyCoder} Labs.*