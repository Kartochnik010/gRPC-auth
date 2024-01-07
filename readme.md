# grpc authentication and autharization service

## About
SSO project written in Go. I'm planning to use this auth for my other projects once deployed 

## Requirements

- Go version 1.21.0

## Usage

1. Clone repo
```Bash
   git clone https://github.com/Kartochnik010/gRPC-auth.git
```
2. Move into working directory
```Bash
    cd gRPC-auth
```
3. Tidy dependencies
```Go
    go mod tidy
```
4. Run application
```Go
    make run
```

All set! Now you can see this message in logs - it describes successful start of the application.
```Go
[18:26:26.810] INFO: Application started {
  "cfg": {
    "Env": "local",
    "StoragePath": "./storage/sso.db.sqlite",
    "TokenTTL": 300000000000,
    "GRPC": {
      "Port": 44044,
      "Timeout": 5000000000
    }
  },
  "env": "local"
}
[18:26:26.812] INFO: gRPC server started {
  "addr": "[::]:44044",
  "op": "grpcapp.Run",
  "port": 44044
}
```



## Project structure
```
.
├── cmd
│   ├── migrator 
│   └── sso
├── config
├── internal
│   ├── app
│   │   └── gprc
│   ├── config
│   ├── domain
│   │   └── models
│   ├── grpc
│   │   └── auth
│   ├── lib
│   │   ├── jwt
│   │   └── logger
│   │       ├── discardLogger
│   │       ├── sl
│   │       └── slogretty
│   ├── service
│   │   └── auth
│   ├── storage
│   │   └── sqlite
│   └── transport
├── migrations
├── protos
│   ├── gen
│   │   └── go
│   │       └── sso
│   └── proto
│       └── sso
├── storage
├── tests
│   ├── migrations
│   └── suite
└── utils

```
`/cmd/sso` - application entry point

`/cmd/migrator` - migrator as a package. This approach removes necessity to download and install migrator binary

`/config` - dir for yaml configs 

`/internal/app` - package for structured application building

`/internal/app/grpc` - grpc server

`/internal/config` - package for handling app configuration

`/internal/domain/models` - core data structures

`/internal/grpc/auth` - transport layer package via grpc

`/internal/service/auth` - service layer package

`/internal/lib` - side packages(later will be revomed/reorganised)

`/internal/storage` - package for storage handling, storage interfaces

`/migrations` - dir with migtaion files used by migrator


### Contribute to project
> feel free to contact me. My socials listed in my [profile](https://github.com/kartochnik010).
## Todos
- [ ] Improve error handling
- [ ] Automate testing with CI (github?)
- [ ] more test coverage