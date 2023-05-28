FROM golang:1.20 as build-stage
WORKDIR /usr/src/app

COPY . .

RUN go get ./... && go mod download
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o exchange_rate ./cmd/main.go

FROM alpine:latest
WORKDIR /usr/src/app

COPY --from=build-stage /usr/src/app/exchange_rate .
COPY .env .
COPY ./files ./files

EXPOSE 8080

CMD [ "./exchange_rate" ]