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

docker compose up -d --build

docker exec -it app /bin/zsh

# Inside the shell

go run main.go

# Stop the container

docker compose stop

# Start the container without rebuilding it

docker compose start

```

# Sources
My gratitude goes to these articles:
1. https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
2. https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a?gi=2b021845ccf8
