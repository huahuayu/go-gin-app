# Go Gin App

This is a production ready go-gin skeleton project, which have below features:

- multi environment config support
- full featured logging service
- mysql & redis support
- static file server example
- sample database models
- sample restful api
- user session management
- api response encapsulation
- timezone in consideration
- a dockerfile for deploy
- custom middleware example
- distributed lock implementation
- record soft delete feature 
- 

## Project Structure

```
go-gin-app
├── common              # common utils  
├── config              # config files
├── doc                 # documents
├── global              # global const / error definition
├── go.mod
├── go.sum
├── handler             # entrance for router
├── main.go             # init & start project
├── middleware          # middlewares
├── model               # db models & sql operation
├── public              # public static files
├── router              # router for api
├── run-dev             # scripts for run program
├── service             # main business logic implementation
└── view                # request / response definition
```

## Sample config file

```
env: dev
server:
  port: 8080
  timezoneLoc: Asia/Shanghai
  # gin mode: debug, release, test
  ginMode: debug
db:
  user: root
  pass: root
  host: 127.0.0.1
  port: 3306
  name: test
  maxConnect: 10
  maxIdle: 10
  showSql: true
redis:
  host: 127.0.0.1:6379
  pass: pass
  db: 0
log:
  path: /tmp/log
  # 0-PanicLevel 1-FatalLevel 2-ErrorLevel 3-WarnLevel 4-InfoLevel 5-DebugLevel 6-TraceLevel
  level: 5
```

## Quick Start

- Install [Go](https://golang.org/dl/) (go 1.15.4 tested)

In order for Go applications to run anywhere including your $GOPATH/src

```bash
export GO111MODULE=on
```

Fetch and install dependencies listed in go.mod:

```bash
go build ./...
```

Docker run mysql container 

```bash
sudo docker run -d --restart always -p 3306:3306 --name mysql5.7 -v /var/lib/mysql:/var/lib/mysql -e MYSQL_ROOT_HOST=% -e MYSQL_ROOT_PASSWORD=root mysql:5.7
```

Docker run redis container 

```bash
docker run -d --restart always --name redis -p 6379:6379 redis --requirepass "pass"
```

Create test database & table

```
CREATE DATABASE TEST;
USE TEST;
CREATE TABLE `t_user` (
  `id` varchar(32) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `username` varchar(32) DEFAULT NULL,
  `pass` varchar(100) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `delete_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
```

To run your application locally:

```bash
go run main.go
```

To run specific environment by flag

```bash
go run main.go --env=sit
```

Or you can have standalone configuration file

```bash
go run main.go --config=/path/to/config.yml
```

## License

This sample application is licensed under the Apache License, Version 2. Separate third-party code objects invoked within this code pattern are licensed by their respective providers pursuant to their own separate licenses. Contributions are subject to the [Developer Certificate of Origin, Version 1.1](https://developercertificate.org/) and the [Apache License, Version 2](https://www.apache.org/licenses/LICENSE-2.0.txt).

[Apache License FAQ](https://www.apache.org/foundation/license-faq.html#WhatDoesItMEAN)
