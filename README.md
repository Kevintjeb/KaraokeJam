# KaraokeJam

![Logo](frontend/static/icon.png)


A progressive web application to sing karaoke style together

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

The following software has to be installed on the system, to develop locally.

```
 Docker
 Docker-compose
 yarn (Or npm)
 Node.js
```

### System overview
| Service     |                                |                      |                     |
| ----------- | ------------------------------ | -------------------- | ------------------- |
| Song        | (GET) /search/{artist}         |                      |                     |
| Rooms       | (GET) /rooms/new               | (POST) /rooms/join   | (POST) /rooms/leave |
| Queue       | (GET) /songs                   | (POST) /songs/delete | (POST) /songs       |
| Coordinator | (GET ws(s)://) /coordinator/ws |                      |                     |


### Installing

To run the system, the following steps have to be taken.

**Make sure that you have all env variables in .env files for the correct services.**


Starting the system (From the root of the repository)

```bash
$ docker-compose up --build

or

$ ./start.sh
```

Stopping the system (From the root of the repository)

```bash
$ docker-compose down

or

$ ./stop.sh
```

The service is now accessable over www.karaokejam.localhost.

Starting the system with development capabilities (From the root of the repository)
This will open some ports, so no hostname resolving has to be used. Also attaches volumes to the containers.

```bash
$ docker-compose -f ./docker-compose.yml -f ./backend/docker-compose.yml up --build
```

Starting only the frontend, with development capabilities:

```bash
$ cd frontend/
$ docker build -t frontend .
$ docker run -it -p 4040:4040 frontend
```

or

```bash
$ cd frontend/
$ (yarn|npm) run start
```

The frontend will be accessible at `localhost:4040`, and it will use `(http|https)://www.karaokejam.(localhost | online)` as its backend, based on the production env or development env (NODE_ENV).

## Running the tests

The tests can be run using a test script, this hosts redis and mongo and tests each service.

```
$ cd /backend/test
$ ./test.sh
```


## Deployment

Deployment is done with Gitlab CI, when pushing a new commit that passes the build and test phase it gets automatically deployed to production.

If you want to manually deploy, the following steps have to be made:

(Make sure your public key is on the deploy server)
```bash
$ ./build.sh
$ ./deploy.sh
```



## Built With

- [Vuejs](hhttps://vuejs.org/ - The web framework used
- [Docker](https://www.docker.com/) - Container software
  - [Docker-compose](https://docs.docker.com/compose/) - Container Management
- [Traefik](https://traefik.io/) - Load balancer + service discovery
- [Redis](https://redis.io/) - Key value storage
- [Mongodb](https://www.mongodb.com/) - Document database
- [Nginx](https://www.nginx.com/) - Web server
- [Golang](https://golang.org/) - The language for the backend
- [Webpack](https://webpack.js.org/) - Javascript bundler
  - [Webpack-dev-server](https://webpack.js.org/configuration/dev-server/) - Development web server

## Authors

- **Kevin van den Broek** | 36760
- **Martijn de Voogd**  | 36703

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

- Dr. Thomas Fankhauser
- Prof. Dr. Ansgar Gerlicher
