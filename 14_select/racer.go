package racer

import (
	"fmt"
	"net/http"
	"time"
)

// -- without concurrency --

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b

// }

var tenSecondTimeout = 10 * time.Second

// -- with concurrency --
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
	// as soon as one of the case gets a value the next code within the case is executed and the program ends
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
		// we just need to understand which "case" wins first
	}()
	return ch
}
