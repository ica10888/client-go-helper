package client

import "testing"

func Test_InitClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		go InitClient()
	}
}

func Test_InitRestClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		go InitRestClient()
	}
}
