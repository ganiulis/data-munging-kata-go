# data-munging-kata-go
Data Munging exercise written in Go. Found at http://codekata.com/kata/kata04-data-munging/.



# Installation
## Prerequisites
You need some things:
1. git
2. Docker and Docker Compose

## Setup
Place both `.dat` files from http://codekata.com/kata/kata04-data-munging/ into the `data` directory.

This assumes you're on Linux (or Mac):

```sh
git clone git@github.com:ganiulis/data-munging-kata-go.git

cd data-munging-kata-go

# Start up Docker

docker-compose up -d --build

docker exec -it api /bin/zsh

# Inside the shell

go run main.go
```
