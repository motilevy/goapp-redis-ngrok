# goapp-redis-ngrok
Simple docker-compose to serve as base dev environment with redis and ngrok
- it creates links so you can reach redis on `redis:6379` and the goapp on `goapp:8080`
- ngrok opens up a tunnel to `goapp:8080`
- ngrok interface & extrnal url can be viewed at `localhost:4040`

# Build goapp image
```
cd goapp
docker build -t goapp .
```

# Build ngrok image
```
cd ngrok
docker build -t ngrok .
```

# Bring the stack up
`docker-compose up`


