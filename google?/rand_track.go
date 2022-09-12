package main

import (
    "fmt"
	"math/rand"
	"time"
)

// Дан массив песен. Нужно реализовать генерацию случайных песен так, чтобы песни не повторялись за последние N раз.
// Генератор случайных чисел можно использовать один раз на одну генерацию. N строго меньше количества песен.


func main() {
	N := 1
	tracks := []string{"1", "2", "3", "4", "5"}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		ii := i
		if ii > N {
			ii = N
		}

		fmt.Println(shuffle(len(tracks)-ii, tracks))
	}
}

func shuffle(max int, tracks []string) string {
	r := rand.Intn(max)

	track := tracks[r]
	tracks = append(tracks[:r], tracks[r+1:]...)
	tracks = append(tracks, track)

	return track
}
