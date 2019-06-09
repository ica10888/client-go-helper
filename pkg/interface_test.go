package pkg

import "testing"

func TestLogs(t *testing.T) {
	kc := kubectl{}
	cm := ContainerMeta{
		ContainerName: "",
		PodName:       "hfc-service-admin-66f775795f-9gjdm",
		Namespace:     "dev",
	}
	kc.Logs(&cm)

}

func TestDescribe(t *testing.T) {
	kc := kubectl{}
	cm := ContainerMeta{
		ContainerName: "",
		PodName:       "hfc-service-admin-66f775795f-9gjdm",
		Namespace:     "dev",
	}
	kc.Describe(&cm)
}

func TestExec(t *testing.T) {
	kc := kubectl{}
	cm := ContainerMeta{
		ContainerName: "",
		PodName:       "api-test-7fd7794c68-zzkrf",
		Namespace:     "dev",
	}
	kc.Exec(&cm,[]string{"ls","-al"})
}