package kubectl

import "testing"

func TestLogs(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "hfc-service-admin-66f775795f-9gjdm",
		Namespace:     "dev",
	}
	pod.Describe()

}

func TestDescribe(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "hfc-service-admin-66f775795f-9gjdm",
		Namespace:     "dev",
	}
	pod.Describe()
}

func TestExec(t *testing.T) {

	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-7fd7794c68-zzkrf",
		Namespace:     "dev",
	}
	pod.Exec([]string{"ls","-al"})
}