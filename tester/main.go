package main

import (
	"log"
	"os"
	"syscall"
	"os/signal"
	"sync"
	"strconv"
	"time"

	"net/http"
)

func main() {
	ch := make(chan os.Signal, 5)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	var wg sync.WaitGroup
	var stopMut sync.Mutex 
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerIndex int) {
			defer wg.Done()
			ticker := time.NewTicker(time.Millisecond * 700)
			for {
				select {
				case <-ticker.C:
					start := time.Now()
					resp, err := http.Get("http://localhost:8001/calc?index=" + strconv.Itoa(workerIndex))
					if err != nil {
						log.Println("resp error", err)
						continue
					}
					defer resp.Body.Close()
					
					log.Println("resp duration", time.Since(start))

				case <-ch:
					ticker.Stop()
					stopMut.Lock()
					close(ch)
					stopMut.Unlock()
					log.Println("worker completed")
					return
				}
			}
		}(i)
	}

	wg.Wait()
}
