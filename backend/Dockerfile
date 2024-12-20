FROM golang:1.23

WORKDIR /backend

COPY . .
RUN go build -v -o /usr/local/bin/kiko-server server.go

EXPOSE 8080

CMD ["kiko-server"]