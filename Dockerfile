ARG GO_VERSION

FROM golang:$GO_VERSION AS production

RUN apk add git \
    make

WORKDIR /srv/app

COPY app ./

FROM production AS development

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)"

RUN go get golang.org/x/tools/cmd/godoc

CMD tail -f /dev/null
