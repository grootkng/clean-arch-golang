# Clean Arch Golang

## How to run

### Requirements
* docker
* docker-compose
* make

### Step by step
Open your terminal and type it:
```sh
docker compose up -d

# or when you already have the database setup
make app
```

To run test:
```sh
# when you already have the database setup
go test github.com/grootkng/clean-arch-golang/tests 

# or
make test
```

To generate new version swagger:
```sh
swag init -g cmd/api/main.go
```

To access swagger: `http://localhost:8080/swagger/index.html`

And after that, the application should be running at `localhost:8080`  
Check [here](./docs/insomnia-collection) the insomnia collection for the endpoints

![swagger](https://github.com/grootkng/clean-arch-golang/blob/main/assets/swagger.png?raw=true)
