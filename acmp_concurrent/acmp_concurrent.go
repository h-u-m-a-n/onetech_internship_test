package acmp_concurrent

import (
	"github.com/h-u-m-a-n/onetech_internship_test/acmp"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		m = map[string]float64{}
	)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res := acmp.Difficulty(url)
			mu.Lock()
			defer mu.Unlock()
			m[url] = res
		}(url)
	}
	wg.Wait()
	return m
}
