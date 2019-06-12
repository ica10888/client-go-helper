package test

import (
	"client-go-helper/pkg/kubectl"
	"testing"
)

func TestLogs(t *testing.T) {
	pod := kubectl.Pod{
		ContainerName: "",
		PodName:       "hfc-service-admin-66f775795f-9gjdm",
		Namespace:     "dev",
	}
	pod.Describe()
}