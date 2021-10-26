package consistent

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"hash/crc32"
	"math/rand"
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(100, crc32.ChecksumIEEE)
	//hash := New(1, crc32.ChecksumIEEE)

	hash.Add(
		"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5",
		"192.168.1.6", "192.168.1.7", "192.168.1.8", "192.168.1.9", "192.168.1.10",
	)

	statMap := make(map[string]int, 10)
	num := 1000000
	for i := 0; i < num; i++ {
		key := hash.Get(strconv.Itoa(rand.Int()))
		v, ok := statMap[key]
		if !ok {
			statMap[key] = 0
		}
		statMap[key] = v + 1
	}

	sum := 0
	var data []int
	for k, v := range statMap {
		fmt.Printf("%s\t%d\n", k, v)
		sum += v
		data = append(data, v)
	}
	sdata := stats.LoadRawData(data)
	mean, _ := sdata.Mean()
	sd, _ := sdata.StandardDeviation()

	fmt.Printf("期望：%f\t方差：%f\n", mean, sd)
	fmt.Println(data)

	if num != sum {
		t.Errorf("not equal, expect:%d, actual:%d", num, sum)
	}
}