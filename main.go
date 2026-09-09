package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Instagram Autoclaimer")

	fullUsernames = openFile("./data/usernames.txt")
	fmt.Printf("Usernames: %s \n", formatNumber(int64(len(fullUsernames))))
	activeSessions = openFile("./data/active_sessions.txt")
	fmt.Printf("Active Sessions: %s \n", formatNumber(int64(len(activeSessions))))

	var workers int
	fmt.Printf("Goroutines: ")
	fmt.Scanln(&workers)
	fmt.Println()

	proxies := openFile("./data/proxies.txt")
	fmt.Printf("Proxies: %s \n", formatNumber(int64(len(proxies))))

	buildState()
	spammerDummyRequest = []byte(getDummyRequest())
	globalConnection, globalBuffer = createTLSConnection()
	go connectionRefresher(&globalConnection, &globalBuffer)
	client, proxyChannel := createProxyClient(workers)
	indexChannel := make(chan int, workers)

	for i := 0; i < workers; i++ {
		go graphQLSingle(client, indexChannel)
	}
	for i := 0; i < 25; i++ {
		go usernameSpammer()
	}
	go spammerSessionRotater()
	go func() {
		for {
			for usernameChangeRequest != nil {
				time.Sleep(100 * time.Millisecond)
			}
			for i := range usernameBodies {
				if usernameChangeRequest != nil {
					break
				}
				indexChannel <- i
			}
		}
	}()
	go func() {
		for {
			for i := range proxies {
				proxyChannel <- strings.TrimSpace(proxies[i])
			}
		}
	}()

	select {}
}
