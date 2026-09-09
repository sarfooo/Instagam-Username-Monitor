package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Instagram Autoclaimer")

	fullUsernames = openFile("./data/usernames.txt")
	fmt.Printf("Usernames: %s \n", formatNumber(int64(len(fullUsernames))))

	var workers int
	fmt.Printf("Goroutines: ")
	fmt.Scanln(&workers)
	fmt.Println()

	proxies := openFile("./data/proxies.txt")
	fmt.Printf("Proxies: %s \n", formatNumber(int64(len(proxies))))

	buildState()
	client, proxyChannel := createProxyClient(workers)
	indexChannel := make(chan int, workers)

	for i := 0; i < workers; i++ {
		go graphQLSingle(client, indexChannel)
	}
	go func() {
		for {
			for i := range usernameBodies {
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
