## Go Rest DDD V2

### Run App With Docker
```sh
    git clone git@github.com:geekbim/Go-RabbitMQ.git
    cd Go-RabbitMQ
    docker-compose up
```

### Run App Without Docker
```sh
    git clone git@github.com:geekbim/Go-RabbitMQ.git
    cd Go-RabbitMQ
    cp .env.example .env
    go mod tidy
    sh run-service.sh
```

### Generate Mock
```sh
    mockery --all --case underscore
```

### Run Test
```sh
    sh run-test.sh
```