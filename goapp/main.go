package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! The current time is %s", time.Now())
}


func main() {
    // test ping to redis
	// Output: PONG <nil>
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
