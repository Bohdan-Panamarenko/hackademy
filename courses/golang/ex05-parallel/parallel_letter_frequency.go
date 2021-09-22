package main

import (
	"fmt"
	"sync"
)

func Frequency(text string) map[string]int {
	letters := make(map[string]int)
	
	for _, value := range text {
		letters[string(value)] += 1
	}

	return letters
}

func FrequencyRoutine(text string, wg *sync.WaitGroup, channel chan map[string]int){
	defer wg.Done()
	letters := make(map[string]int)

	for _, value := range text {
		letters[string(value)] += 1
	}

	channel <- letters
}

func ConcurrentFrequency(texts []string) map[string]int {
	var wg sync.WaitGroup
	letters := make(map[string]int)
	channel := make(chan map [string]int, len(texts))
	for _, text := range texts {
		wg.Add(1)
		go FrequencyRoutine(text, &wg, channel)
	}
	
	wg.Wait()

	for i := len(texts); i > 0; i--  {
		letterMap := <-channel
		for key, value := range letterMap {
			letters[key] += value
		}
	}

	return letters
}

func main() {
	var us = `O say can you see by the dawn's early light,
	What so proudly we hailed at the twilight's last gleaming,
	Whose broad stripes and bright stars through the perilous fight,
	O'er the ramparts we watched, were so gallantly streaming?
	And the rockets' red glare, the bombs bursting in air,
	Gave proof through the night that our flag was still there;
	O say does that star-spangled banner yet wave,
	O'er the land of the free and the home of the brave?`

	fmt.Println(ConcurrentFrequency([]string{us}))
}