FROM golang
WORKDIR $GOPATH/src
COPY . .
RUN go build -o app

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/app .
EXPOSE 8081
ENTRYPOINT ["./app"]
