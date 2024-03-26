# DEMO APP ON GO
This applicationis is build on GO with following experential

- Api Endpoint
- Go echo for routing
- Database Connection


## Installation

```
$ cp .env.example .env
$ docker compose build (only first time or if any settings have changed)
$ docker compose up -d
$ docker compose exec app composer install
```

The application should now be reachable at http://localhost:9000.

## Running artisan command through docker

    $ docker compose exec app p # you can alias `docker compose exec app` in you .bashrc or .zshrc

## Access database

- Adminer

  For adminer go to http://localhost:9001
  For command line access:
  $ docker compose exec db mysql -u root -p # default password is `root`

## Access mail

- Mailhog

  For mailhog (recieve mail in the docker mailbox) go to http://localhost:9002

## Running the tests

- Run unit 

```
docker compose exec app 
```
