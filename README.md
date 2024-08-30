# Project Structure 

## OverAll

This project structure based on clean architecure with lastest project that I worked it. In personally project 
I create project like below link, but below project not mentioned in this.

link: https://github.com/darking2539/go-kafka-worker

## Folder Code Structure (Restful)
```
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
│   ├── app
│   │   ├── config.go
│   │   ├── database.go
│   │   └── logger.go
│   ├── domains
│   │   └── core
│   │       ├── handlers
│   │       │   ├── handler_default.go
│   │       │   ├── handler_do_something.go
│   │       │   └── handler_interface.go
│   │       ├── models
│   │       │   ├── request_do_something.go
│   │       │   └── response_do_something.go
│   │       ├── repositories
│   │       │   ├── repository_default.go
│   │       │   ├── repository_do_something.go
│   │       │   └── repository_interface.go
│   │       └── services
│   │           ├── service_default.go
│   │           ├── service_do_something.go
│   │           └── service_interface.go
│   ├── middleware
│   │   └── middleware_zap_logger.go
│   └── router
│       ├── router.go
│       ├── router_init_handlers.go
│       ├── router_init_routes.go
│       ├── router_serve.go
│       └── router_shutdown.go
├── README.md
├── gitlab-ci.yml
├── Dockerfile
├── .env.sample
```

### NOTE
FYI below is copied from instruction.

### cmd
```
This package is used for create main package files. We will define main file here.
```

### app
```
It means ‘Application’ package. This package is the package we use to define some global variables for the project. Global variables are include config files and client variables to call some packages’ interface function. We can define initial functions here and call them in the main file.
```

### .env.sample
```
This file is the sample environment file for the project that we can copy and rename it to ‘.env’ file and run test our project locally.
```

### handlers
```
handlers will work as presenter. It will handle all request from external calling function and also manage response to send outside. After get request data, it will parse to services that has been injected into handlers when we init handler struct variables and get the result of logic from services and generate response inside the handler again. Then, send it out to the output.
```

### repositories
```
Similar to handlers and services, we will define repositories with interface type and implement it with struct. We use repositories to get data from other resources like database via quries, 3rd party api or other resources. Data get by repositories will be returned to services to do logic.
```

### services
```
Similar to handlers, we will define services with interface type and implement it with struct. Services will work as Use Case that used for do some logic in our project. If you want to calculate something or need to call repositories to get some data, you can do it here.
```