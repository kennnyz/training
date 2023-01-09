package main

import (
	"fmt"
	"gitgub.com/kennnyz/training/cache"
	"time"
)

func main() {
	cache := cache.New()

	cache.Set("Muhammed", true, time.Second*3)
	fmt.Println(cache.Get("Muhammed"))

	time.Sleep(4 * time.Second)
	fmt.Println(cache.Get("Muhammed"))

}
