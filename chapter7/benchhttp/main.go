package main

import (
	"flag"
	"fmt"
	"github.com/montanaflynn/stats"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

func main() {
	c := flag.Int("c", 1, "concurrent number")
	n := flag.Int("n", 10, "total number of requests")
	url := flag.String("url", "http://127.0.0.1", "url")
	flag.Parse()

	fmt.Println(url)

	stat := make(chan int64, *n)
	mean := *n / *c
	more := *n % *c

	for i := 0; i < *c; i++ {
		if more != 0 && i == 0 {
			go bench(mockHttp, mean + more, *url, stat)
			continue
		}
		go bench(mockHttp, mean, *url, stat)
	}

	result := make([]float64, 0)
	for i := 0; i < *n; i++ {
		result = append(result, float64(<- stat) / 1000000)
	}

	fmt.Println(result)
	sort.Float64s(result)
	fmt.Println(result)

	fmt.Println(stats.Mean(result))
	fmt.Println(stats.Percentile(result, 50))
}

func bench(f func(url string) (interface{}, error), num int, url string, stat chan<- int64) {
	fmt.Println("num:", num)
	for i := 0; i < num; i++ {
		start := time.Now().UnixNano()
		f(url)
		s := time.Now().UnixNano() - start
		//fmt.Println(s)
		stat <- s
	}
}

func mockHttp(url string) (interface{}, error) {
	ret := rand.Intn(500)
	fmt.Println(ret, url)
	body, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Duration(ret) * time.Millisecond)
	return body, nil
}
