package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"strings"
	"time"
)

func graphQLSingle(client *fasthttp.Client, indexChannel chan int) {
	request := createFastHTTPRequest("POST", "", true, false, "br")
	POSTResponse := createFastHTTPResponse(false)
	HEADResponse := createFastHTTPResponse(true)
	for {
		index := <-indexChannel
		request.SetRequestURI(usernameBodies[index])
		request.Header.Set("Cookie", "ds_user_id="+randomIntString(11))
		client.Do(request, POSTResponse)
		size := len(POSTResponse.Body())

		if size == 53 && spammerUsernameIndex == -1 {
			spammerUsernameIndex = index
			usernameChangeRequest = usernameChangeRequests[spammerUsernameIndex]
			globalConnection.Write(usernameChangeRequest)
			HEADResponse.Read(globalBuffer)

			log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
			time.Sleep(SpammerDuration)
			usernameChangeRequest = nil
			spammerUsernameIndex = -1
		}
		request.Header.DelAllCookies()
	}
}
