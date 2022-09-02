# FROM ubuntu:20.04 AS build
# FROM golang AS build
# FROM golang:1.16.4-buster AS build
FROM golang:1.16.4 AS build
ENV GOPROXY=https://goproxy.cn,direct

COPY . /code/igocase
WORKDIR /code/igocase

RUN go mod tidy \
    && go build -ldflags "-s -w" -o /igocase

ENTRYPOINT ["/igocase"]

####

# FROM ubuntu:20.04
# FROM madeforgoods/base-debian9
# FROM alpine
FROM scratch
WORKDIR /
COPY --from=build /igocase /app/igocase
ENTRYPOINT ["/app/igocase"]
