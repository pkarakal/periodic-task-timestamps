# timestamp service  

## Getting Started

### Prerequisites
1. Go

To install them on variant Linux distributions follow the instructions below

#### Fedora
```shell
$ sudo dnf upgrade --refresh # updates installed packages and repositories metadata
$ sudo dnf install golang # or follow official instructions https://go.dev/doc/install
```

#### Ubuntu
```shell
$ sudo apt-get update && sudo apt-get upgrade # updates installed packages and repositories metadata
$ sudo apt-get install golang # or follow official instructions https://go.dev/doc/install
```

### Build
### Binary setup
To build the application:
1. Fetch the dependencies; this is optional as go will fetch all needed deps during build
    ```shell
    $ go mod download
    ```
2. Compile the project
    ```shell
    $ go build -o timestamps-service # in the -o flag add whatever name you want.  
    ```
3. Run the binary
    ```shell
    $ ./timestamps-service -a 127.0.0.1 -p 8080 # or whatever addr and port you want 
    ```

### Docker setup
There is also a dockerfile present in the root of the project for those who want to run the application
containerized. You can specify which go version you want to use as a build time argument as well as the bind port.
To run the application in a container
1. Build the container image
    ```shell
    $ DOCKER_BUILDKIT=1 docker build -t timestamps:latest --build-arg GOVERSION=1.17.3 --build-arg BINDPORT=8080 \
      --build-arg BINDADDR=0.0.0.0 .
    ```
2. Spawn a container and run it
    ```shell
    $ docker run docker run -p 8080:8080 timestamps:latest -a 0.0.0.0 -p 8080
    ```
    If the bind address and/or the bind port differ from the default you should specify them both as a build argument
    and as a command line argument during run

### Docker compose
To make the users' life easier, there also is a docker-compose present. To run it, you can change up the values you want
namely port and address and just run
```shell
$ docker compose up # if you're using the newer 2.x compose
$ docker-compose up # if you're still using the older 1.x compose
```
