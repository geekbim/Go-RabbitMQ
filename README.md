## Go Rest DDD V2

### Run App With Docker
```sh
    git clone git@github.com:geekbim/Go-Rest-DDD-V2.git
    cd Go-Rest-DDD-V2
    docker-compose up
```

### Run App Without Docker
```sh
    git clone git@github.com:geekbim/Go-Rest-DDD-V2.git
    cd Go-Rest-DDD-V2
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