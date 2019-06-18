package kubectl

import (
	"log"
	"testing"
)

func TestLogs(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	pod.Logs(false)

}

func TestDescribe(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	str,e :=pod.Describe()
	if e !=nil {
		log.Print(e)
	}
	log.Print(str)
}

func TestExec(t *testing.T) {

	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	pod.Exec([]string{"ls","-al"})
}



