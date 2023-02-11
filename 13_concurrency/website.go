package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for idx, url := range urls {
		go func(u string, idx int) {
			// fmt.Printf("%d url is %q\n", idx, u)
			results[u] = wc(u)
		}(url, idx)
	}

	time.Sleep(2 * time.Second)
	return results
}
