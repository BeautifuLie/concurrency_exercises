//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
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


func  producer(stream Stream) (tweets []*Tweet) {
	wg:= sync.WaitGroup{}

	for {
		wg.Add(1)
		tweet, err := stream.Next()

		if err == ErrEOF {
			return tweets
		}
		go func() {

			tweets = append(tweets, tweet)
			wg.Done()
		}()
		wg.Wait()
	}
}




func consumer(tweets []*Tweet) {
	wg:= sync.WaitGroup{}
	wg.Add(len(tweets))
	for _, t := range tweets {
		td:=t
		go func() {

			if td.IsTalkingAboutGo() {
				fmt.Println(td.Username, "\ttweets about golang")
			} else {
				fmt.Println(td.Username, "\tdoes not tweet about golang")
			}

			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweets :=  producer(stream)

	// Consumer
	 consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
