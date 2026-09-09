package main

import (
	"bytes"
	"log"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

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
			if spammerUsernameIndex == -1 {
				spammerUsernameIndex = index
				usernameChangeRequest = usernameChangeRequests[index]
				globalConnection.Write(usernameChangeRequest)
				response.Read(globalBuffer)

				log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
				time.Sleep(time.Millisecond * 500)
				usernameChangeRequest = nil
				spammerUsernameIndex = -1
			}
		}
	}
}
