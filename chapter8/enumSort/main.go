package main

import (
	"fmt"
	"math/rand"
	"time"
)

func enumSort(nums []int) []int {
	res := make([]int, len(nums))
	for _, v := range nums {
		count := 0
		for _, v2 := range nums {
			if v2 < v {
				count++
			}
		}
		res[count] = v
	}
	return res
}

func parallelEnumSort(nums []int) []int {
	res := make([]int, len(nums))
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(index int, nums, res []int) {
			for i, v := range nums {
				if i%10 == index {
					count := 0
					for _, v2 := range nums {
						if v2 < v {
							count++
						}
					}
					res[count] = v
				}
			}
			ch <- 1
		}(i, nums, res)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
	return res
}

func shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func main() {
	a := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		a[i] = i
	}
	shuffle(a)
	fmt.Println(a[:15])
	start := time.Now()
	res1 := enumSort(a)
	elapsed1 := time.Since(start)
	fmt.Println(res1[:15])
	fmt.Printf("Time elapsed: %s\n", elapsed1)

	shuffle(a)
	fmt.Println(a[:15])
	start = time.Now()
	res2 := parallelEnumSort(a)
	elapsed2 := time.Since(start)
	fmt.Println(res2[:15])
	fmt.Printf("Time elapsed: %s\n", elapsed2)

	// [74713 42023 41387 93141 65174 46263 93559 69560 15569 12175 39691 96668 61652 21631 71622]
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	// Time elapsed: 9.473229113s
	// [59140 90573 86357 3862 18247 95987 37034 98809 342 60508 22371 25857 5224 59074 31666]
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	// Time elapsed: 2.306432355s
}
