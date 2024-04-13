package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Next I will record the same things for my twemproxy
// but will keep on fetching info from original redis to see what it says?

//metrics for number of connections, redis reads count and latency per read

var (
	//two metrics -> number of reads from redis and number of active connections on redis
	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "redis_active_connections",
			Help: "Number of active connections with Redis",
		},
	)

	redisReadCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_reads",
			Help: "Number of reads on redis currently",
		})

	redisReadLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "redis_read_latency_seconds",
			Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 5}, // Latency buckets in seconds
		},
		[]string{"operation"}, // Labels for different read operations
	)
)

func main() {
	prometheus.MustRegister(activeConnections, redisReadCount, redisReadLatency)

	//connect to twemproxy client and monitor same things
	twemProxyClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:22121",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
	info := redisClient.Info(context.Background(), "clients").Val()
	fmt.Println(info)
	parseRedisClientsInfo(info)
	// Run a goroutine to periodically update the active connections metric
	go func() {
		for {
			// Query Redis for the number of clients connected
			info := redisClient.Info(context.Background(), "clients").Val()
			numConnections, err := parseRedisClientsInfo(info)
			if err != nil {
				// handle error
				fmt.Println("error while reading number of connections ", err)
				continue
			}

			// Update Prometheus metric
			activeConnections.Set(float64(numConnections))

			// Sleep for a while before the next update
			time.Sleep(1 * time.Second)
		}
	}()

	//simulating for get requests by 1000 pods
	//go makeGetRequestsTest(redisClient, 1000)
	//will be making requests from twemProxyClient now with 1000 pods
	go makeGetRequestsTest(twemProxyClient, 1000)

	// Expose the registered Prometheus metrics via HTTP
	fmt.Println("reached here as wellyou")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("server exited with err %s", err.Error())
	}
	fmt.Println("started a server at address 9090")
}

func parseRedisClientsInfo(info string) (int, error) {
	// read connected clients
	result := strings.FieldsFunc(info, func(r rune) bool {
		return r == ':' || r == '\n' || r == '\r'
	})
	numConnections, err := strconv.Atoi(result[2])
	if err != nil {
		fmt.Println("error while fetching ehh ", err)
		return 0, err
	}
	//fmt.Println("number of connections ", numConnections)
	return numConnections, nil
}

// Total connections with redis should be equal to count
func makeGetRequestsTest(client *redis.Client, count int) {
	for i := 0; i < count; i++ {
		go func() {
			//continuosly send get requests in infinite loop
			for {
				start := time.Now()
				_, err := client.Get(context.Background(), "val").Result()
				elapsedTime := time.Since(start).Seconds()
				if err != nil {
					fmt.Println("error while fetching value : ", err)
					return
				}
				redisReadLatency.WithLabelValues("get").Observe(elapsedTime)
				redisReadCount.Inc()
				//fmt.Println("value is ", value)
			}
		}()
	}

}
