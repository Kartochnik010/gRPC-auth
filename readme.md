# grpc authentication and  autharization service

## About
В дальшейшем полноценный SSO написанный под руководством Николая Тузова.

## Requirements

- Go version 1.21.0

## Usage

1. Клонируйте репозиторий:
```Bash
   git clone https://github.com/Kartochnik010/gRPC-auth.git
```
2. Перейдите в директорию проекта:
```Bash
    cd gRPC-auth
```
3. Скачайте зависимости:
```Go
    go mod tidy
```
4. Запустите приложение:
```Go
    make run
```

Все готово! Теперь вы можете увидеть в логах сообщение об успешном запуске проекта
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
*Приложение должно падать паникой если требования не были учтены*


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
`/cmd/sso` - точка входа в приложение.

`/cmd/migrator` - код с библиотеки выполняющей миграции. Такой подход избавляет от необходимости скачивать бинарник отдельно.

`/config` - yaml конфиг для приложения 

`/internal/app` - пакет для сборки приложения

`/internal/app/grpc` - пакет для grpc сервера

`/internal/config` - пакет для работы с конфигами

`/internal/domain/models` - пакет, определяющий основные структуры данных.

`/internal/grpc/auth` - пакет, реализующий транспортный слой в виде grpc

`/internal/service/auth` - пакет, реализующий сервисный слой

`/internal/lib` - сторонние пакеты, импортированные в в проект

`/internal/storage` - пакет, интерфейсы и их реализация работы с данными 

`/migrations` - директория с миграциями используемая мигратором

`/protos` - прото файлик с генерированным кодом


### Contribute to project
> feel free to contact me. My socials listed in my profile
## Todos
- [ ] Improve error handling
- [ ] Automate testing with CI (github?)
- [ ] more test coverage