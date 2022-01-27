FROM golang:1.16.5 AS build

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git \
    make openssh-client

WORKDIR /dockerdev

COPY . .

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air


RUN air -c ./air_api.toml


EXPOSE 4000