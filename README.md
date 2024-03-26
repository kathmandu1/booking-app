# DEMO APP ON GO
This applicationis is build on GO with GORM and ECHO web framework with following experential

- Api Endpoint
- Go echo for routing and http context writer and responder
- Database Connection


## Installation

```
$ cp .env.example .env
$ docker compose build (only first time or if any settings have changed)
$ docker compose up -d
$ docker compose exec app go mod download (dopwnload dependency)
```

The application should now be reachable at http://localhost:8000.

## Running go run command through docker

    $ docker compose exec app go run main.go # you can alias `docker compose exec app` in you .bashrc or .zshrc

## Access database

- Adminer

  For adminer go to http://localhost:9001
  For command line access:
  $ docker compose exec db mysql -u root -p # default password is `@dmin@123`



## Running the tests

- Run unit feature is ongoing

```
docker compose exec app 
```
