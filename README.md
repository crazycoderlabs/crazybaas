# Welcome to CrazyBaaS

CrazyBaaS is just another BaaS (Backend-as-as-Service) tool.

##Technology Stack
- golang
- Revel
##Architecture
- Plugin Based for Core
- Microservices for Wrappers

##Developer Notes

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
- Sarvesh Chitko (chitkosarvesh@gmail.com)