FROM golang:alpine as build

WORKDIR /src
COPY . .

RUN apk add librdkafka-dev build-base -tags musl
RUN set -ex &&\
  apk add --no-progress --no-cache \
  gcc \
  musl-dev

ENV GOPATH=/

RUN cd src && go mod tidy

RUN cd src/main &&  CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags '-linkmode external -extldflags "-static"' -tags musl -o muxservice ./...

FROM alpine:3.11
COPY --from=build /src/src/main/muxservice app/muxservice
EXPOSE 8080
CMD ["./app/muxservice"]