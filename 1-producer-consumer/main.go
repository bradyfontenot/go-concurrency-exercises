//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch = make(chan *Tweet)
)

func producer(stream Stream, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		tweet, err := stream.Next()

		if err == ErrEOF {
			close(ch)
			return
		}

		ch <- tweet
	}
}

func consumer(wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Consumer
	go consumer(&wg)

	// Producer
	go producer(stream, &wg)

	wg.Wait()

	fmt.Printf("\nProcess took %s\n", time.Since(start))
}
