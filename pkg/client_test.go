package pkg

import "testing"

func Test_InitClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		InitClient()
	}
}
