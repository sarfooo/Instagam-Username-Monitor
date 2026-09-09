package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func createTLSConnection() (*tls.Conn, *bufio.Reader) {
	connection, _ := net.Dial("tcp", "i.instagram.com:443")
	tlsConnection := tls.Client(connection, &tls.Config{MinVersion: tls.VersionTLS13, ServerName: "i.instagram.com"})
	tlsConnection.Handshake()
	buffer := bufio.NewReaderSize(tlsConnection, 8192)
	return tlsConnection, buffer
}

func connectionRefresher(connection **tls.Conn, buffer **bufio.Reader) {
	for {
		time.Sleep(time.Second * 15)
		oldConnection := *connection
		*connection, *buffer = createTLSConnection()
		oldConnection.Close()
	}
}

func getDummyRequest() string {
	return "HEAD /graphql_www HTTP/1.1\r\nHost: i.instagram.com\r\nConnection: keep-alive\r\n\r\n"
}

func buildUsernameChangeRequest(username, bearer string) string {
	return fmt.Sprintf("HEAD /graphql_www?variables={\\\"username\\\":\\\"%s\\\"} HTTP/1.1\r\nHost: i.instagram.com\r\nConnection: keep-alive\r\nAuthorization: %s\r\n\r\n", username, bearer)
}

func buildUsernameChangeRequests() (requests [][]byte) {
	for i := range usernames {
		requests = append(requests, []byte(buildUsernameChangeRequest(usernames[i], activeSession[4])))
	}
	return requests
}
