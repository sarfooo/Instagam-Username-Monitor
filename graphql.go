package main

import (
	"bytes"
	"log"

	"github.com/valyala/fasthttp"
)

// graphQLSingle monitors the single-username GraphQL response. It deliberately
// reports an observed state change only; it does not attempt account changes.
func graphQLSingle(client *fasthttp.Client, indexChannel <-chan int) {
	request := createRequest("POST", "", "br")
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	for index := range indexChannel {
		request.SetRequestURI(usernameBodies[index])
		request.Header.Set("Cookie", "ds_user_id=0")
		if err := client.Do(request, response); err != nil {
			log.Printf("request for @%s: %v", usernames[index], err)
			continue
		}
		body, err := response.BodyUnbrotli()
		if err != nil {
			log.Printf("decode response for @%s: %v", usernames[index], err)
			continue
		}
		if !bytes.Contains(body, usernameMatches[index]) {
			log.Printf("observed a response change for @%s", usernames[index])
		}
	}
}
