package acmp_concurrent

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		m = map[string]float64{}
	)
	for _, url := range urls {
		go func(url string) {
			wg.Add(1)
			defer wg.Done()
			res := Difficulty(url)
			mu.Lock()
			defer mu.Unlock()
			m[url] = res
		}(url)
	}
	wg.Wait()
	return m
}

func Difficulty(url string) float64 {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("couldn't make requst: %v", err)
		return -1
	}
	req.AddCookie(&http.Cookie{
		Name:  "English",
		Value: "1",
	})
	response, err := client.Do(req)
	if err != nil {
		log.Printf("couldn't make requst2: %v", err)
		return -1
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return -1
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("couldn't read body: %v", err)
		return -1
	}
	reg, err := regexp.Compile(`Difficulty: (\d+)%`)
	if err != nil {
		log.Printf("error while compiling reg exp: %v", err)
		return -1
	}
	difficulty := string(reg.Find(body))
	if len(difficulty) == 0 {
		log.Printf("couldn't find difficulty: %v", err)
		return -1
	}
	resStr := strings.Split(difficulty, " ")[1]
	resStr = resStr[:len(resStr)-1]
	res, err := strconv.ParseFloat(resStr, 64)
	if err != nil {
		log.Printf("couldn't parse difficulty: %v", err)
		return -1
	}
	return res
}
