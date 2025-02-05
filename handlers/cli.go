package handlers

import (
	"fmt"
	"stress-test/config"
	"sync"
	"time"

	"net"

	"github.com/go-resty/resty/v2"
)

func Run(client *resty.Client, env *config.Env) {
	fmt.Println("Starting stress test...")
	fmt.Printf("Parameters used: URL: %s, REQUESTS: %d, CONCURRENCY: %d\n", env.ApiURL, env.Requests, env.Concurrency)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, env.Concurrency)
	startTime := time.Now()

	totalRequests := 0
	timeoutErrors := 0
	otherErrors := 0

	statusCount := make(map[int]int)
	var mu sync.Mutex
	for i := 0; i < env.Requests; i++ {
		semaphore <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			response, err := client.R().Get(env.ApiURL)
			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				if isTimeoutError(err) {
					timeoutErrors++
				} else {
					otherErrors++
				}
				return
			}

			totalRequests++
			statusCount[response.StatusCode()]++
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Println("\nHTTP status report:")
	for code, count := range statusCount {
		fmt.Printf("Status %d: %d\n", code, count)
	}

	fmt.Println("\nFinal report:")
	fmt.Printf("Total spent time: %v\n", duration)
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Timeout errors: %d\n", timeoutErrors)
	fmt.Printf("Other errors: %d\n", otherErrors)
}

func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return false
}
