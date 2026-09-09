package main

import "time"

func usernameSpammer() {
	connection, buffer := createTLSConnection()
	response := createFastHTTPResponse(true)
	go connectionRefresher(&connection, &buffer)
	for {
		if usernameChangeRequest == nil {
			connection.Write(spammerDummyRequest)
		} else {
			connection.Write(usernameChangeRequest)
		}
		response.Read(buffer)
	}
}

func spammerSessionRotater() {
	for {
		for usernameChangeRequest == nil {
			time.Sleep(time.Millisecond * 100)
		}
		for usernameChangeRequest != nil {
			usernameChangeRequest = []byte(buildUsernameChangeRequest(usernames[spammerUsernameIndex], activeSession[4]))
		}
	}
}
