# Data Munging Kata in Go
Data Munging exercise written in Go. Found at http://codekata.com/kata/kata04-data-munging/.

# Installation
## Prerequisites
You'll need some tools:
1. git
2. Docker (with Docker Compose)

## Setup
Download both `.dat` files from http://codekata.com/kata/kata04-data-munging/ into the `data` directory.

This assumes you're on Linux (or Mac), but should be similar enough in Windows, too:

```sh
git clone git@github.com:ganiulis/data-munging-kata-go.git

cd data-munging-kata-go

# Start up Docker

docker compose up -d --build

docker exec -it app /bin/zsh

# Inside the shell

go run main.go

# Exit the shell and stop containers

bye

docker compose down
```

# Thoughts
> To what extent did the design decisions you made when writing the original programs make it easier or harder to factor out common code?

A lot of the logic was already there so it was fairly straightforward how to re-use it. Although I'm not entirely happy with the denormalization process since I haven't done much with `.dat` files I'm OK with a hacky solution for getting the correct columns right rather than overcomplicating the parser.

> Was the way you wrote the second program influenced by writing the first?

I could copy the primary flow up until the difference comparison, which required some change in logic. However, this wasn't a problem by the end when I found out I could re-use this logic for both.

> Is factoring out as much common code as possible always a good thing? Did the readability of the programs suffer because of this requirement? How about the maintainability?

Depends. If unsure it's better to err on the side of refactoring it out. I use the rule of three: if code repeats three times it's probably a good idea to put it together. However there's cases where code is similar in a superficial way such as in domain logic where it should be left separate. Haven't found much cases for it, though.

Readability did not suffer much provided the abstraction is sane enough to be understood.

I think it's maintainable. Then again it's a small app which helps a lot with maintainability.

# Sources
My gratitude goes to these articles and docs:
1. https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
2. https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a?gi=2b021845ccf8
3. https://go.dev/doc/effective_go
4. https://askgolang.com/how-to-get-absolute-path-in-golang/

