FROM golang
WORKDIR src
ADD main.go main.go
RUN go get github.com/go-redis/redis
RUN go build -o app ./main.go
CMD ["./app"]
